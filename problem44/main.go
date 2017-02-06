package main

import (
	"fmt"
	"math"
)

func pent(i int) int {
	return (3*i*i - i) >> 1
}

func invPent(x int) (int, bool) {
	var i int
	f := (1 + math.Sqrt(1+24*float64(x))) / 6
	i = int(f)
	return i, pent(i) == x
}

func main() {
	for i, dp, p := 1, 1, 1; i < 1<<32; i, dp, p = i+1, dp+3, p+dp+3 {
		for j, dq, q := 1, 1, 1; dq <= p; j, dq, q = j+1, dq+3, q+dq+3 {
			if _, ok := invPent(q + p); ok {
				if _, ok := invPent(q + q + p); ok {
					fmt.Println(p)
					return
				}
			}
		}
	}
}
