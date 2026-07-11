package proof

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type fileUploader interface {
	UploadFile(context.Context, string) (string, error)
}

type linkMatch struct {
	start int
	end   int
	path  string
}

var (
	inlineLinkPattern = regexp.MustCompile(`!?\[[^\]\r\n]*\]\(\s*(?:<([^>\r\n]+)>|([^\s)]+))(?:\s+(?:"[^"]*"|'[^']*'|\([^)]*\)))?\s*\)`)
	bareLinkPattern   = regexp.MustCompile(`\[((?:\.\.?/)[^\]\r\n]+)\]`)
)

// PrepareDocument replaces relative Markdown link targets with uploaded signed URLs.
func PrepareDocument(ctx context.Context, uploader fileUploader, documentPath string, source []byte) ([]byte, error) {
	matches := findRelativeLinks(source)
	if len(matches) == 0 {
		return source, nil
	}

	baseDir := filepath.Dir(documentPath)
	resolved := make(map[string]string, len(matches))
	for _, match := range matches {
		parsed, err := url.Parse(match.path)
		if err != nil {
			return nil, fmt.Errorf("parse relative link %q: %w", match.path, err)
		}
		localPath := filepath.Clean(filepath.Join(baseDir, filepath.FromSlash(parsed.Path)))
		info, err := os.Stat(localPath)
		if err != nil {
			return nil, fmt.Errorf("validate relative link %q: %w", match.path, err)
		}
		if !info.Mode().IsRegular() {
			return nil, fmt.Errorf("validate relative link %q: not a regular file", match.path)
		}
		resolved[match.path] = localPath
	}

	urls := make(map[string]string, len(resolved))
	for relativePath, localPath := range resolved {
		signedURL, err := uploader.UploadFile(ctx, localPath)
		if err != nil {
			return nil, fmt.Errorf("upload relative link %q: %w", relativePath, err)
		}
		urls[relativePath] = signedURL
	}

	result := append([]byte(nil), source...)
	sort.Slice(matches, func(i, j int) bool { return matches[i].start > matches[j].start })
	for _, match := range matches {
		result = append(result[:match.start], append([]byte(urls[match.path]), result[match.end:]...)...)
	}
	return result, nil
}

func findRelativeLinks(source []byte) []linkMatch {
	matches := make([]linkMatch, 0)
	occupied := make([][2]int, 0)
	for _, indexes := range inlineLinkPattern.FindAllSubmatchIndex(source, -1) {
		start, end := indexes[2], indexes[3]
		if start < 0 {
			start, end = indexes[4], indexes[5]
		}
		path := string(source[start:end])
		if isRelativeFileLink(path) {
			matches = append(matches, linkMatch{start: start, end: end, path: path})
		}
		occupied = append(occupied, [2]int{indexes[0], indexes[1]})
	}
	for _, indexes := range bareLinkPattern.FindAllSubmatchIndex(source, -1) {
		if overlaps(indexes[0], indexes[1], occupied) {
			continue
		}
		matches = append(matches, linkMatch{start: indexes[2], end: indexes[3], path: string(source[indexes[2]:indexes[3]])})
	}
	return matches
}

func isRelativeFileLink(value string) bool {
	parsed, err := url.Parse(strings.TrimSpace(value))
	if err != nil || parsed.Scheme != "" || parsed.Host != "" || parsed.Path == "" || strings.HasPrefix(parsed.Path, "/") {
		return false
	}
	return !strings.HasPrefix(value, "#")
}

func overlaps(start, end int, ranges [][2]int) bool {
	for _, candidate := range ranges {
		if start < candidate[1] && end > candidate[0] {
			return true
		}
	}
	return false
}
