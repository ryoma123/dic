package dic

import (
	"testing"
)

func TestToUint16(t *testing.T) {
	// https://github.com/miekg/dns/blob/master/types.go
	m := map[string]uint16{
		"a":     1,
		"ns":    2,
		"cname": 5,
		"soa":   6,
		"ptr":   12,
		"mx":    15,
		"txt":   16,
		"aaaa":  28,
		"any":   255,
	}

	for k := range m {
		if toUint16(k) != m[k] {
			t.Errorf("Failed toUint16: %s", k)
		}
	}
}
