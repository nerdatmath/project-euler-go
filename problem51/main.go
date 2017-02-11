package main

import (
	"fmt"
	"strconv"

	"github.com/nerdatmath/project-euler-go/primes"
)

// "51404" -> {"*1404", "5*404", "51*04", "51*0*", "514*4"}
func families(s string) []string {
	l := []string{}
	for _, c := range "0123456789" {
		var f func(string, string)
		f = func(s, prefix string) {
			for i, cc := range s {
				if cc == c {
					ss := s[:i] + "*" + s[i+1:]
					l = append(l, prefix+ss)
					f(ss[i:], prefix+ss[:i])
				}
			}
		}
		f(s, "")
	}
	return l
}

func main() {
	m := make(map[string][]string)
	for _, p := range primes.Primes {
		s := strconv.Itoa(p)
		for _, family := range families(s) {
			l, _ := m[family]
			l = append(l, s)
			m[family] = l
			if len(l) >= 8 {
				// fmt.Println(family, l)
				fmt.Println(l[0])
				return
			}
		}
	}
}
