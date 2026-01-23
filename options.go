package dic

const defaultCnameMax = 5

// Options controls lookup behavior.
type Options struct {
	Reverse     bool
	FollowCNAME bool
	CnameMax    int
}

func normalizeOptions(o Options) Options {
	if o.CnameMax <= 0 {
		o.CnameMax = defaultCnameMax
	}
	return o
}
