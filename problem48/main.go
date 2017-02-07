package main

import "fmt"

func powmod(b, e, m int) int {
	p := 1
	for ; e > 0; e >>= 1 {
		if e&1 == 1 {
			p = (p * b) % m
		}
		b = (b * b) % m
	}
	return p
}

func main() {
	m := int(1e10)
	m2 := 1 << 10
	m5 := m / m2
	invm2m5 := 1745224 // modular inverse of 2^10 modulo 5^10
	invm5m2 := 841     // modular inverse of 5^10 modulo 2^10

	sum2, sum5 := 0, 0
	for n := 1; n <= 1000; n++ {
		sum2 = (sum2 + powmod(n, n, m2)) % m2
		sum5 = (sum5 + powmod(n, n, m5)) % m5
	}
	sum := (sum2*invm5m2*m5 + sum5*invm2m5*m2) % m
	fmt.Println(sum)
}
