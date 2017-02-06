package main

import "fmt"

var primes [1<<16 + 3]int

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
}

func factors4(x int) bool {
	factors := 0
	for _, p := range primes {
		if x == 1 || factors >= 4 {
			return x == 1 && factors == 4
		}
		if p == 0 || x%p != 0 {
			continue
		}
		factors++
		for x%p == 0 {
			x /= p
		}
	}
	panic("not enough primes")
}

func main() {
	count := 0
	for i := 1; ; i++ {
		if !factors4(i) {
			count = 0
			continue
		}
		count++
		if count == 4 {
			fmt.Println(i - 3)
			break
		}
	}
}
