package dic

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

// Dic struct
type Dic struct {
	domains   []string
	servers   []string
	nsRecords map[string]string
}

// New returns Dic
func New(d []string) *Dic {
	var s []string
	var ns map[string]string

	for _, a := range confSec {
		s = append(s, a.Server)

		if strings.EqualFold(a.Server, "ns") && len(ns) == 0 {
			ns = setNSRecords(d)
		}
	}
	wg.Wait()

	return &Dic{
		domains:   d,
		servers:   s,
		nsRecords: ns,
	}
}

func setNSRecords(domains []string) map[string]string {
	m := map[string]string{}

	for _, d := range domains {
		wg.Add(1)

		go func(d string) {
			defer wg.Done()
			r := query(d, qtypes["ns"], resolv.Servers[0])

			// In the case of CNAME, the result output of "NS" is skipped
			if len(r.Answer) != 0 {
				if _, ok := r.Answer[0].(*dns.CNAME); ok {
					w := fmt.Sprintf("[%s] CNAME record is set, so do not query NS record", d)
					setNotice(w)
					return
				}
			}

			if len(r.Answer) == 0 {
				r = query(getSuperDomain(d), qtypes["ns"], resolv.Servers[0])

				// If the super domain NS record is not found either, the result output of "NS" is skipped
				if len(r.Answer) == 0 {
					w := fmt.Sprintf("[%s %s] NS records not found. Please check if the domain is valid", d, getSuperDomain(d))
					setNotice(w)
					return
				}
			}

			if ns, ok := r.Answer[0].(*dns.NS); ok {
				w := strings.Fields(ns.String())

				mu.Lock()
				m[d] = w[4]
				mu.Unlock()
			}
		}(d)
	}

	return m
}

func getSuperDomain(d string) string {
	s := strings.Split(d, ".")
	return strings.Join(s[1:], ".")
}

func setServer(s string, d string, dic *Dic) string {
	switch strings.ToLower(s) {
	case "":
		return resolv.Servers[0]
	case "ns":
		return dic.nsRecords[d]
	default:
		return s
	}
}
