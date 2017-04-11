package num

func PowMod(b, e, m int) int {
	p := 1
	for ; e > 0; e >>= 1 {
		if e&1 == 1 {
			p = (p * b) % m
		}
		b = (b * b) % m
	}
	return p
}
