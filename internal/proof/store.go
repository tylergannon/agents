package proof

import (
	"context"
	"fmt"
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

// SignedURLLifetime is R2's maximum supported S3 presigned URL lifetime.
const SignedURLLifetime = 7 * 24 * time.Hour

type s3API interface {
	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	ListObjectsV2(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	DeleteObjects(context.Context, *s3.DeleteObjectsInput, ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
}

type presignAPI interface {
	PresignGetObject(context.Context, *s3.GetObjectInput, ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

// Store uploads, signs, and removes proof artifacts in an S3-compatible bucket.
type Store struct {
	cfg     Config
	client  s3API
	presign presignAPI
	now     func() time.Time
}

// NewStore creates a Cloudflare R2-backed store.
func NewStore(ctx context.Context, cfg Config) (*Store, error) {
	awsCfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion("auto"),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("load S3 configuration: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(options *s3.Options) {
		options.BaseEndpoint = aws.String(cfg.Endpoint)
		options.UsePathStyle = true
	})
	return &Store{cfg: cfg, client: client, presign: s3.NewPresignClient(client), now: time.Now}, nil
}

// UploadFile uploads path and returns its configured public or signed download URL.
func (s *Store) UploadFile(ctx context.Context, path string) (signedURL string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open %q: %w", path, err)
	}
	defer func() {
		if closeErr := file.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close %q: %w", path, closeErr)
		}
	}()

	info, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("stat %q: %w", path, err)
	}
	if !info.Mode().IsRegular() {
		return "", fmt.Errorf("%q is not a regular file", path)
	}

	key := s.objectKey(filepath.Base(path))
	input := &s3.PutObjectInput{Bucket: aws.String(s.cfg.Bucket), Key: aws.String(key), Body: file}
	if contentType := mime.TypeByExtension(filepath.Ext(path)); contentType != "" {
		input.ContentType = aws.String(contentType)
	}
	if _, err := s.client.PutObject(ctx, input); err != nil {
		return "", fmt.Errorf("upload %q: %w", path, err)
	}

	return s.downloadURL(ctx, key)
}

func (s *Store) objectKey(name string) string {
	cleanName := strings.ReplaceAll(filepath.Base(name), " ", "-")
	return fmt.Sprintf("%s/%s/%s-%s", s.cfg.Prefix, s.now().UTC().Format("2006/01/02"), uuid.NewString(), cleanName)
}

func (s *Store) downloadURL(ctx context.Context, key string) (string, error) {
	if s.cfg.DownloadMode == DownloadModePublic {
		return s.publicURL(key), nil
	}
	request, err := s.presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.cfg.Bucket),
		Key:    aws.String(key),
	}, func(options *s3.PresignOptions) {
		options.Expires = SignedURLLifetime
	})
	if err != nil {
		return "", fmt.Errorf("sign %q: %w", key, err)
	}
	return request.URL, nil
}

func (s *Store) publicURL(key string) string {
	parts := strings.Split(key, "/")
	for index, part := range parts {
		parts[index] = url.PathEscape(part)
	}
	return s.cfg.PublicBaseURL + "/" + strings.Join(parts, "/")
}

// Vacuum deletes objects under the configured prefix older than maxAge.
func (s *Store) Vacuum(ctx context.Context, maxAge time.Duration) (int, error) {
	cutoff := s.now().Add(-maxAge)
	var continuation *string
	deleted := 0
	for {
		page, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket:            aws.String(s.cfg.Bucket),
			Prefix:            aws.String(s.cfg.Prefix + "/"),
			ContinuationToken: continuation,
		})
		if err != nil {
			return deleted, fmt.Errorf("list objects: %w", err)
		}

		objects := make([]types.ObjectIdentifier, 0, len(page.Contents))
		for _, object := range page.Contents {
			if object.LastModified != nil && object.LastModified.Before(cutoff) && object.Key != nil {
				objects = append(objects, types.ObjectIdentifier{Key: object.Key})
			}
		}
		if len(objects) > 0 {
			_, err = s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
				Bucket: aws.String(s.cfg.Bucket),
				Delete: &types.Delete{Objects: objects, Quiet: aws.Bool(true)},
			})
			if err != nil {
				return deleted, fmt.Errorf("delete objects: %w", err)
			}
			deleted += len(objects)
		}

		if !aws.ToBool(page.IsTruncated) || page.NextContinuationToken == nil {
			return deleted, nil
		}
		continuation = page.NextContinuationToken
	}
}
