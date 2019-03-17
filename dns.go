package dic

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/miekg/dns"
)

const timeout = 5

var qtypes = map[string]uint16{
	"a":     dns.TypeA,
	"ns":    dns.TypeNS,
	"cname": dns.TypeCNAME,
	"soa":   dns.TypeSOA,
	"ptr":   dns.TypePTR,
	"mx":    dns.TypeMX,
	"txt":   dns.TypeTXT,
	"aaaa":  dns.TypeAAAA,
	"any":   dns.TypeANY,
}

var resolv *dns.ClientConfig

func setResolv() {
	c, err := dns.ClientConfigFromFile(resolvFile)
	if err != nil {
		err := fmt.Errorf(err.Error())
		setError(err)
	}
	resolv = c
}

func query(d string, q uint16, s string) *dns.Msg {
	c := new(dns.Client)
	c.Dialer = &net.Dialer{
		Timeout: time.Duration(timeout) * time.Second,
	}

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(d), q)
	m.RecursionDesired = true

	r, _, err := c.Exchange(m, net.JoinHostPort(s, resolv.Port))
	if r == nil {
		err := fmt.Errorf(err.Error())
		setError(err)
	}
	return r
}

// toUint16 to uint16 for string
func toUint16(q string) uint16 {
	return qtypes[strings.ToLower(q)]
}
