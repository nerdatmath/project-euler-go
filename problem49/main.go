package main

import (
	"fmt"
	"sort"

	"github.com/nerdatmath/project-euler-go/primes"
)

func sortDigits(x int) int {
	digits := []int{}
	for ; x > 0; x /= 10 {
		digits = append(digits, x%10)
	}
	sort.Ints(digits)
	out := 0
	for _, digit := range digits {
		out = out*10 + digit
	}
	return out
}

func triples(xs ...int) [][3]int {
	if len(xs) < 3 {
		return nil
	}
	var out [][3]int
	for i, a := range xs[:len(xs)-2] {
		for j, b := range xs[i+1 : len(xs)-1] {
			for _, c := range xs[i+j+2:] {
				out = append(out, [3]int{a, b, c})
			}
		}
	}
	return out
}

func main() {
	a := sort.SearchInts(primes.Primes, 1000)
	b := sort.SearchInts(primes.Primes, 10000)
	groups := make(map[int][]int)
	for _, p := range primes.Primes[a:b] {
		k := sortDigits(p)
		g, _ := groups[k]
		groups[k] = append(g, p)
	}
	for _, ps := range groups {
		for _, t := range triples(ps...) {
			if t[0] == 1487 && t[1] == 4817 && t[2] == 8147 {
				continue
			}
			if t[1]-t[0] == t[2]-t[1] {
				fmt.Printf("%d%d%d\n", t[0], t[1], t[2])
			}
		}
	}
}
