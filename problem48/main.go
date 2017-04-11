package main

import (
	"fmt"

	"github.com/nerdatmath/project-euler-go/num"
)

func main() {
	m := int(1e10)
	m2 := 1 << 10
	m5 := m / m2
	invm2m5 := 1745224 // modular inverse of 2^10 modulo 5^10
	invm5m2 := 841     // modular inverse of 5^10 modulo 2^10

	sum2, sum5 := 0, 0
	for n := 1; n <= 1000; n++ {
		sum2 = (sum2 + num.PowMod(n, n, m2)) % m2
		sum5 = (sum5 + num.PowMod(n, n, m5)) % m5
	}
	sum := (sum2*invm5m2*m5 + sum5*invm2m5*m2) % m
	fmt.Println(sum)
}
