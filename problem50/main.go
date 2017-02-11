package main

import (
	"fmt"

	"github.com/nerdatmath/project-euler-go/primes"
)

func main() {
	if primes.Primes[len(primes.Primes)-1] < 1e6 {
		panic(primes.Primes[len(primes.Primes)-1])
	}
	var sum int
	var sums []int
	for _, p := range primes.Primes {
		sum += p
		sums = append(sums, sum)
		if sum > 2e6 {
			break
		}
	}
	for l := len(sums); ; l-- {
		for i := 0; i < len(sums)-l; i++ {
			p := sums[i+l-1] - sums[i]
			if p > 1e6 {
				continue
			}
			if primes.IsPrime(p) {
				fmt.Println(p)
				return
			}
		}
	}
}
