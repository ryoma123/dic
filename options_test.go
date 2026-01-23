package dic

import "testing"

func TestNormalizeOptions(t *testing.T) {
	opts := normalizeOptions(Options{CnameMax: 0})
	if opts.CnameMax != defaultCnameMax {
		t.Errorf("Expected default CnameMax %d, got %d", defaultCnameMax, opts.CnameMax)
	}

	opts = normalizeOptions(Options{CnameMax: 3})
	if opts.CnameMax != 3 {
		t.Errorf("Expected CnameMax 3, got %d", opts.CnameMax)
	}
}
