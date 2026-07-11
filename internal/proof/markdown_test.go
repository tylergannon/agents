package proof

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fakeUploader struct {
	paths []string
}

func (f *fakeUploader) UploadFile(_ context.Context, path string) (string, error) {
	f.paths = append(f.paths, path)
	return "https://signed.example/" + filepath.Base(path), nil
}

func TestPrepareDocumentRewritesRelativeLinks(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	for _, name := range []string{"foo.jpg", "report.pdf"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(name), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	source := []byte("[./foo.jpg]\n![screenshot](./foo.jpg)\n[report](report.pdf)\n[web](https://example.com)\n")
	uploader := &fakeUploader{}

	got, err := PrepareDocument(context.Background(), uploader, filepath.Join(dir, "proof.md"), source)
	if err != nil {
		t.Fatal(err)
	}
	want := "[https://signed.example/foo.jpg]\n![screenshot](https://signed.example/foo.jpg)\n[report](https://signed.example/report.pdf)\n[web](https://example.com)\n"
	if string(got) != want {
		t.Fatalf("prepared document:\n%s\nwant:\n%s", got, want)
	}
	if len(uploader.paths) != 2 {
		t.Fatalf("uploaded %d files, want 2", len(uploader.paths))
	}
}

func TestPrepareDocumentValidatesBeforeUploading(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	uploader := &fakeUploader{}
	_, err := PrepareDocument(context.Background(), uploader, filepath.Join(dir, "proof.md"), []byte("[missing](./missing.png)"))
	if err == nil || !strings.Contains(err.Error(), "validate relative link") {
		t.Fatalf("error = %v", err)
	}
	if len(uploader.paths) != 0 {
		t.Fatalf("uploaded %d files before validation completed", len(uploader.paths))
	}
}
