package main

import "fmt"
import "github.com/nerdatmath/project-euler-go/primes"

func factors4(x int) bool {
	factors := 0
	for _, p := range primes.Primes {
		if x == 1 || factors >= 4 {
			return x == 1 && factors == 4
		}
		if x%p != 0 {
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
