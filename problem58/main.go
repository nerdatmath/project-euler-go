package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := int64(1)
	n := int64(0)
	for i := int64(3); ; i += 2 {
		for j := 0; j < 4; j++ {
			x += i - 1
			if big.NewInt(int64(x)).ProbablyPrime(50) {
				n++
			}
		}
		if n*10 < i*2-1 {
			fmt.Println(i)
			return
		}
	}
}
