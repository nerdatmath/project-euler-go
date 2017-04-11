package num

import "math/big"

var ten = big.NewInt(10)

func Reverse(x *big.Int) *big.Int {
	z := new(big.Int).Set(x)
	y := new(big.Int)
	m := new(big.Int)
	for z.Sign() != 0 {
		z.DivMod(z, ten, m)
		y.Mul(y, ten).Add(y, m)
	}
	return y
}

func Palindrome(x *big.Int) bool {
	return x.Cmp(Reverse(x)) == 0
}
