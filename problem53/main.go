package main

import "fmt"

func main() {
	target := int(1e6)
	count := 0
	n, r, nCr := 0, 0, 1
	for n <= 100 {
		// loop invariant: nCr <= target
		// loop invariant: 0 <= r <= n - r <= n <= 100
		if r < n-r-1 {
			// Increase r
			nCr *= n - r
			r++
			nCr /= r
		}
		if nCr > target {
			// We just stepped over the boundary
			// nCx > target for all r <= x <= n-r
			count += n - r - r + 1
		}
		// Increase n
		n++
		nCr *= n
		nCr /= n - r
		// Decrease r while nCr > target
		for nCr > target {
			nCr *= r
			r--
			nCr /= n - r
		}
	}
	fmt.Println(count)
}
