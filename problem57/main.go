package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewRat(1, 1)
	one := big.NewRat(1, 1)
	count := 0
	for i := 0; i < 1000; i++ {
		x.Add(x, one)
		x.Inv(x)
		x.Add(x, one)
		if len(x.Num().Text(10)) > len(x.Denom().Text(10)) {
			count++
		}
	}
	fmt.Println(count)
}
