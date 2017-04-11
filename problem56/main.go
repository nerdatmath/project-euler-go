package main

import (
	"fmt"
	"math/big"
)

func main() {
	max := 0
	for a := big.NewInt(1); a.Int64() < 100; a.Add(a, big.NewInt(1)) {
		for b := big.NewInt(0); b.Int64() < 100; b.Add(b, big.NewInt(1)) {
			x := new(big.Int).Exp(a, b, nil)
			sum := 0
			for x.Sign() > 0 {
				_, m := x.DivMod(x, big.NewInt(10), new(big.Int))
				sum += int(m.Int64())
			}
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}
