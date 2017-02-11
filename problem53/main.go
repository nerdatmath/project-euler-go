package main

import "fmt"

func main() {
	target := int(1e6)
	count := 0
	nCr := 1
	n, r, p := 0, 0, 0
	for n < 100 && nCr <= target {
		n++
		nCr *= n
		if r < p {
			r++
			nCr /= r
		} else {
			p++
			nCr /= p
		}
	}
	// We have nCr > target now, and n is the minimum that allows this
	for n <= 100 {
		for nCr > target && r > 1 {
			// reduce r until nCr <= target
			nCr *= r
			r--
			p++
			nCr /= p
		}
		if nCr <= target {
			// back up one
			nCr *= p
			p--
			r++
			nCr /= r
		}
		fmt.Println(n, r)
		// nCx > target for all r <= x <= n-r
		count += n - r - r + 1
		n++
		nCr *= n
		r++
		nCr /= r
	}
	fmt.Println(count)
}
