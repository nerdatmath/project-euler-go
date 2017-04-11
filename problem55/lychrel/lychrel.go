package lychrel

import (
	"math/big"

	"github.com/nerdatmath/project-euler-go/num"
)

func Lychrel(x int) bool {
	z := big.NewInt(int64(x))
	for i := 0; i < 50; i++ {
		z.Add(z, num.Reverse(z))
		if num.Palindrome(z) {
			return false
		}
	}
	return true
}
