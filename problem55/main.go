package main

import (
	"fmt"

	"github.com/nerdatmath/project-euler-go/problem55/lychrel"
)

func main() {
	count := 0
	for i := 1; i < 10000; i++ {
		if lychrel.Lychrel(i) {
			count++
		}
	}
	fmt.Println(count)
}
