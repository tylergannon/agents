package proof

import (
	"context"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type fakeS3 struct {
	listed      []types.Object
	deleted     []types.ObjectIdentifier
	uploadedKey string
	uploaded    string
	contentType string
}

func (f *fakeS3) PutObject(_ context.Context, input *s3.PutObjectInput, _ ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	contents, err := io.ReadAll(input.Body)
	if err != nil {
		return nil, err
	}
	f.uploadedKey = aws.ToString(input.Key)
	f.uploaded = string(contents)
	f.contentType = aws.ToString(input.ContentType)
	return &s3.PutObjectOutput{}, nil
}

func (f *fakeS3) ListObjectsV2(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return &s3.ListObjectsV2Output{Contents: f.listed}, nil
}

type fakePresigner struct {
	key     string
	expires time.Duration
}

func (f *fakePresigner) PresignGetObject(_ context.Context, input *s3.GetObjectInput, options ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	presignOptions := s3.PresignOptions{}
	for _, option := range options {
		option(&presignOptions)
	}
	f.key = aws.ToString(input.Key)
	f.expires = presignOptions.Expires
	return &v4.PresignedHTTPRequest{URL: "https://signed.example/download"}, nil
}

func TestUploadFileSendsContentsAndReturnsConfiguredURL(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "screen shot.png")
	if err := os.WriteFile(path, []byte("image bytes"), 0o644); err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 7, 11, 12, 0, 0, 0, time.UTC)
	client := &fakeS3{}
	presigner := &fakePresigner{}
	store := &Store{
		cfg:     Config{Bucket: "proofs", Prefix: "proof", DownloadMode: DownloadModeSigned},
		client:  client,
		presign: presigner,
		now:     func() time.Time { return now },
	}

	got, err := store.UploadFile(context.Background(), path)
	if err != nil {
		t.Fatal(err)
	}
	if got != "https://signed.example/download" {
		t.Fatalf("signed URL = %q", got)
	}
	if client.uploaded != "image bytes" || client.contentType != "image/png" {
		t.Fatalf("uploaded contents = %q, content type = %q", client.uploaded, client.contentType)
	}
	if !strings.HasPrefix(client.uploadedKey, "proof/2026/07/11/") || !strings.HasSuffix(client.uploadedKey, "-screen-shot.png") {
		t.Fatalf("uploaded key = %q", client.uploadedKey)
	}
	if presigner.key != client.uploadedKey || presigner.expires != SignedURLLifetime {
		t.Fatalf("presigned key = %q, expires = %v", presigner.key, presigner.expires)
	}
}

func (f *fakeS3) DeleteObjects(_ context.Context, input *s3.DeleteObjectsInput, _ ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	f.deleted = append(f.deleted, input.Delete.Objects...)
	return &s3.DeleteObjectsOutput{}, nil
}

func TestVacuumDeletesOnlyExpiredObjects(t *testing.T) {
	t.Parallel()
	now := time.Date(2026, 7, 11, 12, 0, 0, 0, time.UTC)
	client := &fakeS3{listed: []types.Object{
		{Key: aws.String("proof/old"), LastModified: aws.Time(now.Add(-2 * time.Hour))},
		{Key: aws.String("proof/new"), LastModified: aws.Time(now.Add(-30 * time.Minute))},
	}}
	store := &Store{cfg: Config{Bucket: "proofs", Prefix: "proof"}, client: client, now: func() time.Time { return now }}

	count, err := store.Vacuum(context.Background(), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 || len(client.deleted) != 1 || aws.ToString(client.deleted[0].Key) != "proof/old" {
		t.Fatalf("deleted = %#v, count = %d", client.deleted, count)
	}
}

func TestSignedDownloadURLUsesR2MaximumLifetime(t *testing.T) {
	t.Parallel()
	store, err := NewStore(context.Background(), Config{
		AccountID:       "account",
		AccessKeyID:     "access",
		SecretAccessKey: "secret",
		Bucket:          "proofs",
		Endpoint:        "https://account.r2.cloudflarestorage.com",
		Prefix:          "proof",
		DownloadMode:    DownloadModeSigned,
	})
	if err != nil {
		t.Fatal(err)
	}
	signed, err := store.downloadURL(context.Background(), "proof/example.txt")
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := url.Parse(signed)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := parsed.Query().Get("X-Amz-Expires"), "604800"; got != want {
		t.Fatalf("X-Amz-Expires = %q, want %q", got, want)
	}
}

func TestPublicDownloadURLUsesConfiguredBaseURL(t *testing.T) {
	t.Parallel()
	store := &Store{
		cfg: Config{
			DownloadMode:  DownloadModePublic,
			PublicBaseURL: "https://proof.example.r2.dev",
		},
	}

	public, err := store.downloadURL(context.Background(), "proof/a folder/example image.png")
	if err != nil {
		t.Fatal(err)
	}
	if want := "https://proof.example.r2.dev/proof/a%20folder/example%20image.png"; public != want {
		t.Fatalf("public URL = %q, want %q", public, want)
	}
}
