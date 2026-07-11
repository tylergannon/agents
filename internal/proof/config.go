package proof

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	DownloadModePublic = "public"
	DownloadModeSigned = "signed"
)

// Config contains the Cloudflare R2 connection and download settings.
type Config struct {
	AccountID       string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Endpoint        string
	Prefix          string
	DownloadMode    string
	PublicBaseURL   string
}

// LoadConfig loads ~/.proof-uploader/config.yaml, then .env and environment overrides.
func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	v := viper.New()
	v.SetConfigFile(configPath())
	v.SetDefault("prefix", "proof")
	v.SetDefault("download_mode", DownloadModePublic)
	for key, environment := range map[string]string{
		"account_id":        "R2_ACCOUNT_ID",
		"access_key_id":     "R2_ACCESS_KEY_ID",
		"secret_access_key": "R2_SECRET_ACCESS_KEY",
		"bucket":            "R2_BUCKET",
		"endpoint":          "R2_ENDPOINT",
		"prefix":            "R2_PREFIX",
		"download_mode":     "PROOF_DOWNLOAD_MODE",
		"public_base_url":   "PROOF_PUBLIC_BASE_URL",
	} {
		if err := v.BindEnv(key, environment); err != nil {
			return Config{}, fmt.Errorf("bind %s: %w", environment, err)
		}
	}
	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) && !os.IsNotExist(err) {
			return Config{}, fmt.Errorf("read config %q: %w", v.ConfigFileUsed(), err)
		}
	}

	cfg := Config{
		AccountID:       strings.TrimSpace(v.GetString("account_id")),
		AccessKeyID:     strings.TrimSpace(v.GetString("access_key_id")),
		SecretAccessKey: strings.TrimSpace(v.GetString("secret_access_key")),
		Bucket:          strings.TrimSpace(v.GetString("bucket")),
		Endpoint:        strings.TrimRight(strings.TrimSpace(v.GetString("endpoint")), "/"),
		Prefix:          strings.Trim(strings.TrimSpace(v.GetString("prefix")), "/"),
		DownloadMode:    strings.ToLower(strings.TrimSpace(v.GetString("download_mode"))),
		PublicBaseURL:   strings.TrimRight(strings.TrimSpace(v.GetString("public_base_url")), "/"),
	}
	if cfg.Endpoint == "" && cfg.AccountID != "" {
		cfg.Endpoint = "https://" + cfg.AccountID + ".r2.cloudflarestorage.com"
	}

	missing := make([]string, 0, 5)
	for name, value := range map[string]string{
		"account_id":        cfg.AccountID,
		"access_key_id":     cfg.AccessKeyID,
		"secret_access_key": cfg.SecretAccessKey,
		"bucket":            cfg.Bucket,
		"endpoint":          cfg.Endpoint,
	} {
		if value == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return Config{}, fmt.Errorf("missing required configuration: %s", strings.Join(missing, ", "))
	}
	if cfg.DownloadMode != DownloadModePublic && cfg.DownloadMode != DownloadModeSigned {
		return Config{}, fmt.Errorf("download_mode must be %q or %q", DownloadModePublic, DownloadModeSigned)
	}
	if cfg.DownloadMode == DownloadModePublic && cfg.PublicBaseURL == "" {
		return Config{}, fmt.Errorf("public_base_url is required when download_mode is %q", DownloadModePublic)
	}

	return cfg, nil
}

func configPath() string {
	if configured := strings.TrimSpace(os.Getenv("PROOF_UPLOADER_CONFIG")); configured != "" {
		return configured
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".proof-uploader", "config.yaml")
	}
	return filepath.Join(home, ".proof-uploader", "config.yaml")
}
