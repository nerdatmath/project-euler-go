package main

import (
	"fmt"
	"sort"
)

func digits(x int) int {
	ds := make([]int, 0, 15)
	for x > 0 {
		ds = append(ds, x%10)
		x /= 10
	}
	sort.Ints(ds)
	for _, d := range ds {
		x *= 10
		x += d
	}
	return x
}

func main() {
outer:
	for i := 1; ; i++ {
		d := digits(i)
		for m := 2; m <= 6; m++ {
			if d != digits(m*i) {
				continue outer
			}
		}
		fmt.Println(i)
		return
	}
}
