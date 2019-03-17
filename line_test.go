package dic

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	l := lines{
		line{1, 0, 0, "", "", "", nil},
		line{0, 1, 0, "", "", "", nil},
		line{1, 0, 1, "", "", "", nil},
		line{0, 0, 1, "", "", "", nil},
		line{1, 1, 0, "", "", "", nil},
		line{0, 0, 0, "", "", "", nil},
	}
	e := lines{
		line{0, 0, 0, "", "", "", nil},
		line{0, 0, 1, "", "", "", nil},
		line{0, 1, 0, "", "", "", nil},
		line{1, 0, 0, "", "", "", nil},
		line{1, 0, 1, "", "", "", nil},
		line{1, 1, 0, "", "", "", nil},
	}

	if l.sort(); !reflect.DeepEqual(l, e) {
		t.Errorf("Failed sort")
	}
}

func TestGetServerLine(t *testing.T) {
	s := "dns.example.jp"
	m := map[string]string{
		"":               "  -",
		"ns":             " *@dns.example.jp",
		"dns.example.jp": "  @dns.example.jp",
	}

	if e := ""; getServerLine(e, s) != m[e] {
		t.Errorf("Failed getServerLine: %s", e)
	}
	if e := "ns"; getServerLine(e, s) != m[e] {
		t.Errorf("Failed getServerLine: %s", e)
	}
	if e := "dns.example.jp"; getServerLine(e, s) != m[e] {
		t.Errorf("Failed getServerLine: %s", e)
	}
}
