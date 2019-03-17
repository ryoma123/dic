package dic

import (
	"testing"
)

func TestGetSuperDomain(t *testing.T) {
	s := "www.example.com"

	if e := "example.com"; getSuperDomain(s) != e {
		t.Errorf("Failed getSuperDomain: %s", e)
	}
}
