package proof

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigReadsYAMLAndEnvironmentOverrides(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "config.yaml")
	config := []byte(`
account_id: account
access_key_id: access
secret_access_key: secret
bucket: from-file
prefix: artifacts
public_base_url: https://public.example.test
`)
	if err := os.WriteFile(configPath, config, 0o600); err != nil {
		t.Fatal(err)
	}
	clearConfigEnvironment(t)
	t.Setenv("PROOF_UPLOADER_CONFIG", configPath)
	t.Setenv("R2_BUCKET", "from-environment")

	got, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got.Bucket != "from-environment" || got.DownloadMode != DownloadModePublic {
		t.Fatalf("config = %#v", got)
	}
	if got.Endpoint != "https://account.r2.cloudflarestorage.com" {
		t.Fatalf("endpoint = %q", got.Endpoint)
	}
}

func TestLoadConfigAllowsSignedModeWithoutPublicURL(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "config.yaml")
	config := []byte(`
account_id: account
access_key_id: access
secret_access_key: secret
bucket: proofs
download_mode: signed
`)
	if err := os.WriteFile(configPath, config, 0o600); err != nil {
		t.Fatal(err)
	}
	clearConfigEnvironment(t)
	t.Setenv("PROOF_UPLOADER_CONFIG", configPath)

	got, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got.DownloadMode != DownloadModeSigned || got.PublicBaseURL != "" {
		t.Fatalf("config = %#v", got)
	}
}

func clearConfigEnvironment(t *testing.T) {
	t.Helper()
	for _, name := range []string{
		"R2_ACCOUNT_ID",
		"R2_ACCESS_KEY_ID",
		"R2_SECRET_ACCESS_KEY",
		"R2_BUCKET",
		"R2_ENDPOINT",
		"R2_PREFIX",
		"PROOF_DOWNLOAD_MODE",
		"PROOF_PUBLIC_BASE_URL",
	} {
		t.Setenv(name, "")
	}
}
