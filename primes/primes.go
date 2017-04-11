package primes

import "sort"

var primes [1<<24 + 3]int

var Primes []int

func init() {
	copy(primes[:3], []int{2, 3, 5})
	wheel := [...]int{7, 11, 13, 17, 19, 23, 29, 31}
	for p, offset := primes[3:], 0; len(p) >= len(wheel); p, offset = p[len(wheel):], offset+30 {
		for j, n := range wheel {
			p[j] = offset + n
		}
	}
	for i, p := range primes {
		for _, d := range primes[:i] {
			if d == 0 {
				continue
			}
			if p%d == 0 {
				primes[i] = 0
				break
			}
			if d*d > p {
				break
			}
		}
	}
	j := 0
	for _, p := range primes {
		if p == 0 {
			continue
		}
		primes[j] = p
		j++
	}
	Primes = primes[:j]
}

func IsPrime(x int) bool {
	n := sort.SearchInts(Primes, x)
	return n < len(Primes) && Primes[n] == x
}
