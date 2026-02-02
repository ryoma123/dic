package dic

import (
	"bytes"
	"os"
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

func getAppPath(s string) string {
	switch s {
	case configFile:
		if len(configPath) != 0 {
			return configPath
		}
	case sectionFile:
		if len(sectionPath) != 0 {
			return sectionPath
		}
	}

	if exists(s) {
		return s
	}

	return os.ExpandEnv(stringsJoin([]string{appPath, s}))
}
