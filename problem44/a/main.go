package main

import (
	"fmt"
    "math"
)

func isPentagonal(x int) bool {
	var h float64
	h = (math.Sqrt(float64(24*x+1)) + 1) / 6
	return math.Floor(h) == h
}

func pent(n int) int {
    return n*(3*n-1) / 2
}

func main() {
	for n := 2; ; n++ {
		for m := 1; m < n; m++ {
			diff := pent(n) - pent(m)
			if isPentagonal(diff) {
				if isPentagonal(pent(n) + pent(m)) {
					fmt.Println(diff)
					return
				}
			}
		}
	}
}
