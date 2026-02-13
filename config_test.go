package dic

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigDirFromEnv(t *testing.T) {
	dir := t.TempDir()
	os.Setenv("DIC_CONFIG_DIR", dir)
	defer os.Unsetenv("DIC_CONFIG_DIR")

	got := configDir()
	if got != dir {
		t.Errorf("configDir() = %q, want %q", got, dir)
	}
}

func TestConfigDirFromXDG(t *testing.T) {
	os.Unsetenv("DIC_CONFIG_DIR")
	dir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", dir)
	defer os.Unsetenv("XDG_CONFIG_HOME")

	got := configDir()
	want := filepath.Join(dir, "dic")
	if got != want {
		t.Errorf("configDir() = %q, want %q", got, want)
	}
}

func TestConfigDirDefault(t *testing.T) {
	os.Unsetenv("DIC_CONFIG_DIR")
	os.Unsetenv("XDG_CONFIG_HOME")

	got := configDir()
	home, _ := os.UserHomeDir()
	want := filepath.Join(home, ".config", "dic")
	if got != want {
		t.Errorf("configDir() = %q, want %q", got, want)
	}
}

func TestNewConfigDefault(t *testing.T) {
	dir := t.TempDir()
	os.Setenv("DIC_CONFIG_DIR", dir)
	defer os.Unsetenv("DIC_CONFIG_DIR")

	// config.toml が無い状態でデフォルト設定が読み込まれる
	c := newConfig()
	if len(c.Sec) != 1 {
		t.Fatalf("expected 1 section, got %d", len(c.Sec))
	}
	if c.Sec[0].Name != "default" {
		t.Errorf("section name = %q, want %q", c.Sec[0].Name, "default")
	}
	if len(c.Sec[0].Args) != 1 {
		t.Fatalf("expected 1 args, got %d", len(c.Sec[0].Args))
	}
	if c.Sec[0].Args[0].Server != "" {
		t.Errorf("server = %q, want empty", c.Sec[0].Args[0].Server)
	}
}

func TestGetDefaultSectionFallback(t *testing.T) {
	dir := t.TempDir()
	os.Setenv("DIC_CONFIG_DIR", dir)
	defer os.Unsetenv("DIC_CONFIG_DIR")

	// .section ファイルが無い状態で "default" が返る
	got := getDefaultSection()
	if got != "default" {
		t.Errorf("getDefaultSection() = %q, want %q", got, "default")
	}
}

func TestGetDefaultSectionFromFile(t *testing.T) {
	dir := t.TempDir()
	os.Setenv("DIC_CONFIG_DIR", dir)
	defer os.Unsetenv("DIC_CONFIG_DIR")

	// .section ファイルに "custom" を書き込む
	err := os.WriteFile(filepath.Join(dir, sectionFile), []byte("custom\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	got := getDefaultSection()
	if got != "custom" {
		t.Errorf("getDefaultSection() = %q, want %q", got, "custom")
	}
}
