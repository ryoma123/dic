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

	for di, d := range dic.domains {
		for ai, a := range confSec {
			s := setServer(a.Server, d, dic)
			if len(s) == 0 {
				continue
			}

			for qi, q := range a.Qtypes {
				wg.Add(1)
				go func(l line) {
					defer wg.Done()

					r := query(l.domain, toUint16(l.qtype), s)

					for _, a := range r.Answer {
						l.answer = a
						ls = append(ls, l)
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
