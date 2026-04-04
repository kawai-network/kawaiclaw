package internal

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sipeed/picoclaw/pkg/config"
)

func TestGetConfigPath(t *testing.T) {
	// In development mode, paths.Base() returns "data"
	got := GetConfigPath()
	want := filepath.Join("data", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithKAWAICLAW_HOME(t *testing.T) {
	t.Setenv(config.EnvHome, "/custom/kawaiclaw")

	got := GetConfigPath()
	want := filepath.Join("/custom/kawaiclaw", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithKAWAICLAW_CONFIG(t *testing.T) {
	t.Setenv(config.EnvConfig, "/custom/config.json")
	t.Setenv(config.EnvHome, "/custom/kawaiclaw")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := "/custom/config.json"

	assert.Equal(t, want, got)
}

func TestGetConfigPath_Windows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-specific HOME behavior varies; run on windows")
	}

	testUserProfilePath := `C:\Users\Test`
	t.Setenv("USERPROFILE", testUserProfilePath)

	got := GetConfigPath()
	want := filepath.Join(testUserProfilePath, ".kawaiclaw", "config.json")

	require.True(t, strings.EqualFold(got, want), "GetConfigPath() = %q, want %q", got, want)
}
