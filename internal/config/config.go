// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
)

// DefaultBaseURL is the production Addigy API host. The OpenAPI spec only
// declares the relative server path "/api/v2" (no host), so without this
// default a freshly-configured CLI cannot reach the API even with a valid
// key. Override with ADDIGY_BASE_URL or base_url in the config file (e.g.
// for region-specific or test hosts).
const DefaultBaseURL = "https://api.addigy.com"

type Config struct {
	BaseURL                   string            `toml:"base_url,omitempty"`
	BasePath                  string            `toml:"base_path"`
	AuthHeaderVal             string            `toml:"auth_header"`
	Headers                   map[string]string `toml:"headers,omitempty"`
	AuthSource                string            `toml:"-"`
	AccessToken               string            `toml:"access_token"`
	RefreshToken              string            `toml:"refresh_token"`
	TokenExpiry               time.Time         `toml:"token_expiry"`
	ClientID                  string            `toml:"client_id"`
	ClientSecret              string            `toml:"client_secret"`
	Path                      string            `toml:"-"`
	AddigyDocumentationApiKey string            `toml:"documentation_api_key"`
}

func Load(configPath string) (*Config, error) {
	cfg := &Config{
		BaseURL:  DefaultBaseURL,
		BasePath: "/api/v2",
	}

	// Resolve config path
	path := configPath
	if path == "" {
		path = os.Getenv("ADDIGY_CONFIG")
	}
	if path == "" {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, ".config", "addigy-cli", "config.toml")
	}
	cfg.Path = path

	// Try to load config file
	data, err := os.ReadFile(path)
	if err == nil {
		if err := toml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parsing config %s: %w", path, err)
		}
	}

	// Env var overrides
	if v := os.Getenv("ADDIGY_DOCUMENTATION_API_KEY"); v != "" {
		cfg.AddigyDocumentationApiKey = v
		cfg.AuthSource = "env:ADDIGY_DOCUMENTATION_API_KEY"
	}

	// Label config-file-derived credentials so doctor can distinguish
	// "credentials persisted on disk" from "no credentials at all" — without
	// this, users who saved via set-token without an env var see a blank
	// auth_source and can't tell whether their config is being picked up.
	// The label is the literal "config" rather than "config:<path>"; the
	// config file path is exposed separately as report["config_path"], and
	// embedding it in auth_source leaks the user's home directory through
	// doctor's JSON envelope.
	if cfg.AuthSource == "" && (cfg.AuthHeaderVal != "" || cfg.AccessToken != "") {
		cfg.AuthSource = "config"
	}
	if cfg.AuthSource == "" && cfg.AddigyDocumentationApiKey != "" {
		cfg.AuthSource = "config"
	}

	// Base URL override (used by verify to point at mock/test servers)
	if v := os.Getenv("ADDIGY_BASE_URL"); v != "" {
		cfg.BaseURL = v
	}
	if v := os.Getenv("ADDIGY_BASE_PATH"); v != "" {
		cfg.BasePath = v
	}

	// Final fallback: a config file written by an older build (or with an
	// explicit empty value) can blank BaseURL during Unmarshal, so re-assert
	// the production default when nothing else set it. Without this, such a
	// config produces hostless "/api/v2/..." request URLs that never connect.
	if strings.TrimSpace(cfg.BaseURL) == "" {
		cfg.BaseURL = DefaultBaseURL
	}
	return cfg, nil
}

func (c *Config) AuthHeader() string {
	if c.AuthHeaderVal != "" {
		return c.AuthHeaderVal
	}
	return c.AddigyDocumentationApiKey
}

func (c *Config) SaveTokens(clientID, clientSecret, accessToken, refreshToken string, expiry time.Time) error {
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	c.AccessToken = accessToken
	c.RefreshToken = refreshToken
	c.TokenExpiry = expiry
	return c.save()
}

// SaveCredential persists a single API credential to the field that
// AuthHeader() consults for api_key auth. Writing to AccessToken (the
// bearer slot) would silently no-op since AuthHeader() reads the env-var-
// derived field, not AccessToken, when Auth.Type == "api_key".
//
// The clears precede the assignment so a canonical env-var whose placeholder
// collides with a builtin tag (e.g. an env var named XXX_ACCESS_TOKEN
// resolving to the AccessToken field) ends up holding the new token.
func (c *Config) SaveCredential(token string) error {
	c.AuthHeaderVal = ""
	c.AccessToken = ""
	c.AddigyDocumentationApiKey = token
	return c.save()
}

func (c *Config) ClearTokens() error {
	c.AccessToken = ""
	c.RefreshToken = ""
	c.TokenExpiry = time.Time{}
	return c.save()
}

func (c *Config) save() error {
	dir := filepath.Dir(c.Path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}
	data, err := toml.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshaling config: %w", err)
	}
	return os.WriteFile(c.Path, data, 0o600)
}
