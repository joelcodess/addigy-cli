// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package config

import (
	"os"
	"path/filepath"
	"testing"
)

// clearBaseEnv removes the base-URL env overrides so each test controls them
// explicitly (t.Setenv restores prior values after the test).
func clearBaseEnv(t *testing.T) {
	t.Helper()
	t.Setenv("ADDIGY_BASE_URL", "")
	t.Setenv("ADDIGY_BASE_PATH", "")
	t.Setenv("ADDIGY_DOCUMENTATION_API_KEY", "")
}

func TestLoad_DefaultBaseURL_NoFile(t *testing.T) {
	clearBaseEnv(t)
	// Point at a path that does not exist so no file is loaded.
	missing := filepath.Join(t.TempDir(), "nope.toml")
	cfg, err := Load(missing)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.BaseURL != DefaultBaseURL {
		t.Errorf("BaseURL = %q, want default %q", cfg.BaseURL, DefaultBaseURL)
	}
	if cfg.BasePath != "/api/v2" {
		t.Errorf("BasePath = %q, want /api/v2", cfg.BasePath)
	}
}

func TestLoad_EmptyBaseURLInFileFallsBackToDefault(t *testing.T) {
	clearBaseEnv(t)
	// A config written by an older build can persist base_url = "" explicitly;
	// Load must re-assert the default rather than produce a hostless URL.
	path := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(path, []byte("base_url = ''\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.BaseURL != DefaultBaseURL {
		t.Errorf("BaseURL = %q, want default %q (empty file value must fall back)", cfg.BaseURL, DefaultBaseURL)
	}
}

func TestLoad_FileBaseURLIsUsed(t *testing.T) {
	clearBaseEnv(t)
	path := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(path, []byte("base_url = 'https://self-hosted.example.com'\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.BaseURL != "https://self-hosted.example.com" {
		t.Errorf("BaseURL = %q, want the file value", cfg.BaseURL)
	}
}

func TestLoad_EnvOverridesFileAndDefault(t *testing.T) {
	clearBaseEnv(t)
	path := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(path, []byte("base_url = 'https://from-file.example.com'\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("ADDIGY_BASE_URL", "https://from-env.example.com")
	t.Setenv("ADDIGY_BASE_PATH", "/api/v9")
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.BaseURL != "https://from-env.example.com" {
		t.Errorf("BaseURL = %q, env must win over file", cfg.BaseURL)
	}
	if cfg.BasePath != "/api/v9" {
		t.Errorf("BasePath = %q, env must win", cfg.BasePath)
	}
}

func TestLoad_APIKeyEnvSetsAuthSource(t *testing.T) {
	clearBaseEnv(t)
	t.Setenv("ADDIGY_DOCUMENTATION_API_KEY", "test-key-value")
	cfg, err := Load(filepath.Join(t.TempDir(), "nope.toml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.AuthHeader() != "test-key-value" {
		t.Errorf("AuthHeader() = %q, want the env key", cfg.AuthHeader())
	}
	if cfg.AuthSource != "env:ADDIGY_DOCUMENTATION_API_KEY" {
		t.Errorf("AuthSource = %q, want env source", cfg.AuthSource)
	}
}
