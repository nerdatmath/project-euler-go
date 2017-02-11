package main

import "fmt"

func main() {
	target := int(1e6)
	count := 0
	n, r, nCr := 0, 0, 1
	for n <= 100 {
		// loop invariant: nCr == count(n, r)
		// loop invariant: nCr <= target
		// loop invariant: 0 <= r <= n/2
		// loop invariant: r (and therefore nCr) are maximized
		// invariant: nCr > target or r == n/2
		if n-r-r-1 > 0 {
			// nCx > target for all r+1 <= x <= n-r-1
			count += n - r - r - 1
		}
		// Increase n, which may take nCr over the target
		n++
		nCr *= n
		nCr /= n - r
		if r < n-r-1 && nCr <= target {
			// Increase r
			nCr *= n - r
			r++
			nCr /= r
		}
		// Decrease r while nCr > target
		for nCr > target {
			nCr *= r
			r--
			nCr /= n - r
		}
	}
	fmt.Println(count)
}
