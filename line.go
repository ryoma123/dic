package dic

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/miekg/dns"
)

const (
	space           = " "
	unspecified     = "-"
	nsRecordsPrefix = "*@"
	serverPrefix    = "@"
	extraQIndexA    = 1000
	extraQIndexAAAA = 1001
	extraQIndexPTR  = 1002
)

var wg sync.WaitGroup
var mu sync.Mutex

// lines struct
type lines []line

// line struct
type line struct {
	dIndex int
	sIndex int
	qIndex int
	domain string
	server string
	qtype  string
	answer dns.RR
}

// newLines returns lines
func newLines(dic *Dic) lines {
	var ls lines
	var lsMu sync.Mutex

	for di, d := range dic.domains {
		isReverseIP := dic.opts.Reverse && isIPAddress(d)
		for ai, a := range confSec {
			s := setServer(a.Server, d, dic)
			if len(s) == 0 {
				continue
			}

			if isReverseIP {
				wg.Add(1)
				go func(l line, ip string, server string) {
					defer wg.Done()

					rname, ok := reverseAddr(ip)
					if !ok {
						return
					}

					r := query(rname, qtypes["ptr"], server)
					appendAnswers(&ls, &lsMu, l, r.Answer, "ptr", extraQIndexPTR)
				}(line{di, ai, 0, d, s, "ptr", nil}, d, s)
				continue
			}

			for qi, q := range a.Qtypes {
				wg.Add(1)
				go func(l line) {
					defer wg.Done()

					r := query(l.domain, toUint16(l.qtype), s)

					appendAnswers(&ls, &lsMu, l, r.Answer, l.qtype, l.qIndex)

					if dic.opts.FollowCNAME {
						if target, ok := cnameTarget(r.Answer); ok && !hasAddress(r.Answer) {
							followCNAME(&ls, &lsMu, l, target, s, dic.opts)
						}
					}
				}(line{di, ai, qi, d, s, q, nil})
			}
		}
	}
	wg.Wait()

	ls.sort()

	return ls
}

func (l lines) output(dic *Dic) {
	di, si := -1, -1

	for _, l := range l {
		if di != l.dIndex {
			if di != -1 {
				fmt.Print("\n")
			}
			fmt.Printf("[%s]\n", dic.domains[l.dIndex])
		}

		if si != l.sIndex || (di != l.dIndex && si == l.sIndex) {
			fmt.Printf("%s\n", getServerLine(dic.servers[l.sIndex], l.server))
		}

		fmt.Printf("%s%s\n", strings.Repeat(space, 4), l.answer)
		di, si = l.dIndex, l.sIndex
	}
	fmt.Print("\n")
}

func (l lines) sort() {
	sort.SliceStable(l, func(i, j int) bool { return l[i].qIndex < l[j].qIndex })
	sort.SliceStable(l, func(i, j int) bool { return l[i].sIndex < l[j].sIndex })
	sort.SliceStable(l, func(i, j int) bool { return l[i].dIndex < l[j].dIndex })
}

func getServerLine(confServer string, server string) string {
	var s []string

	switch strings.ToLower(confServer) {
	case "":
		s = []string{strings.Repeat(space, 2), unspecified}
	case "ns":
		s = []string{space, nsRecordsPrefix, server}
	default:
		s = []string{strings.Repeat(space, 2), serverPrefix, server}
	}

	return stringsJoin(s)
}

func appendAnswers(ls *lines, mu *sync.Mutex, l line, answers []dns.RR, qtype string, qIndex int) []string {
	var addrs []string

	for _, a := range answers {
		l.answer = a
		l.qtype = qtype
		l.qIndex = qIndex

		mu.Lock()
		*ls = append(*ls, l)
		mu.Unlock()

		switch rr := a.(type) {
		case *dns.A:
			addrs = append(addrs, rr.A.String())
		case *dns.AAAA:
			addrs = append(addrs, rr.AAAA.String())
		}
	}

	return addrs
}

func cnameTarget(answers []dns.RR) (string, bool) {
	for _, a := range answers {
		if rr, ok := a.(*dns.CNAME); ok {
			return rr.Target, true
		}
	}
	return "", false
}

func hasAddress(answers []dns.RR) bool {
	for _, a := range answers {
		switch a.(type) {
		case *dns.A, *dns.AAAA:
			return true
		}
	}
	return false
}

func followCNAME(ls *lines, mu *sync.Mutex, base line, target string, server string, opts Options) {
	name := target

	for depth := 0; depth < opts.CnameMax; depth++ {
		hasAddr := false
		nextTarget := ""

		rA := query(name, qtypes["a"], server)
		addrsA := appendAnswers(ls, mu, base, rA.Answer, "a", extraQIndexA)
		if len(addrsA) != 0 {
			hasAddr = true
		}
		if nextTarget == "" {
			if t, ok := cnameTarget(rA.Answer); ok {
				nextTarget = t
			}
		}
		if opts.Reverse {
			appendPTR(ls, mu, base, addrsA, server)
		}

		rAAAA := query(name, qtypes["aaaa"], server)
		addrsAAAA := appendAnswers(ls, mu, base, rAAAA.Answer, "aaaa", extraQIndexAAAA)
		if len(addrsAAAA) != 0 {
			hasAddr = true
		}
		if nextTarget == "" {
			if t, ok := cnameTarget(rAAAA.Answer); ok {
				nextTarget = t
			}
		}
		if opts.Reverse {
			appendPTR(ls, mu, base, addrsAAAA, server)
		}

		if nextTarget == "" || hasAddr {
			return
		}
		name = nextTarget
	}
}

func appendPTR(ls *lines, mu *sync.Mutex, base line, addrs []string, server string) {
	for _, ip := range addrs {
		rname, ok := reverseAddr(ip)
		if !ok {
			continue
		}
		r := query(rname, qtypes["ptr"], server)
		appendAnswers(ls, mu, base, r.Answer, "ptr", extraQIndexPTR)
	}
}
