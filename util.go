package dic

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getDuplicate(s []string) string {
	m := map[string]string{}

	for _, v := range s {
		v = strings.ToLower(v)

		if _, ok := m[v]; ok {
			return m[v]
		}
		m[v] = v
	}
	return ""
}

func stringsJoin(s []string) string {
	var buf bytes.Buffer

	for _, v := range s {
		buf.WriteString(v)
	}
	return buf.String()
}

func removeNewline(s string) string {
	return strings.NewReplacer(
		"\r\n", "",
		"\r", "",
		"\n", "",
	).Replace(s)
}

// overrideConfigPath is set by --config flag via setConfigPath.
var overrideConfigPath string

func configDir() string {
	if overrideConfigPath != "" {
		return filepath.Dir(overrideConfigPath)
	}
	if v := os.Getenv("DIC_CONFIG_DIR"); v != "" {
		return v
	}
	if v := os.Getenv("XDG_CONFIG_HOME"); v != "" {
		return filepath.Join(v, "dic")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		setError(fmt.Errorf("failed to get home directory: %s", err))
	}
	return filepath.Join(home, ".config", "dic")
}

func configFilePath(name string) string {
	if overrideConfigPath != "" && name == configFile {
		return overrideConfigPath
	}
	return filepath.Join(configDir(), name)
}

func ensureConfigDir() error {
	return os.MkdirAll(configDir(), 0755)
}

func setConfigPath(path string) {
	if len(path) == 0 {
		overrideConfigPath = ""
		return
	}
	overrideConfigPath = path
}
