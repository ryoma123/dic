package dic

import (
	"testing"
)

func TestContains(t *testing.T) {
	s := []string{"sec1", "sec2", "sec3"}

	if e := "sec1"; !contains(s, e) {
		t.Errorf("Failed contains: %s", e)
	}
	if e := "sec4"; contains(s, e) {
		t.Errorf("Failed contains: %s", e)
	}
}

func TestExists(t *testing.T) {
	if e := "util_test.go"; !exists(e) {
		t.Errorf("Failed exists: %s", e)
	}
	if e := "notExist.go"; exists(e) {
		t.Errorf("Failed exists: %s", e)
	}
}

func TestGetDuplicate(t *testing.T) {
	s := []string{"sec1", "sec1", "sec3"}

	if e := "sec1"; getDuplicate(s) != e {
		t.Errorf("Failed getDuplicate: %s", e)
	}
}

func TestStringsJoin(t *testing.T) {
	s := []string{"str1", "str2", "str3"}

	if e := "str1str2str3"; stringsJoin(s) != e {
		t.Errorf("Failed stringsJoin: %s", e)
	}
}

func TestRemoveNewline(t *testing.T) {
	s := "s\r\ne\rc\n"

	if e := "sec"; removeNewline(s) != e {
		t.Errorf("Failed removeNewline: %s", e)
	}
}
