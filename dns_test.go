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

func TestIsIPAddress(t *testing.T) {
	if !isIPAddress("1.2.3.4") {
		t.Errorf("Expected IPv4 to be detected")
	}
	if !isIPAddress("2001:db8::1") {
		t.Errorf("Expected IPv6 to be detected")
	}
	if isIPAddress("example.com") {
		t.Errorf("Did not expect hostname to be detected as IP")
	}
}

func TestReverseAddr(t *testing.T) {
	r, ok := reverseAddr("1.2.3.4")
	if !ok {
		t.Fatalf("Expected reverseAddr to succeed")
	}
	if e := "4.3.2.1.in-addr.arpa."; r != e {
		t.Errorf("Expected %q, got %q", e, r)
	}

	if _, ok := reverseAddr("not-an-ip"); ok {
		t.Errorf("Expected reverseAddr to fail for invalid IP")
	}
}
