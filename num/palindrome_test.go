package num

import (
	"math/big"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		fwd int64
		rev int64
	}{
		{1, 1},
		{10, 1},
		{101, 101},
		{196, 691},
		{98, 89},
		{1000, 1},
		{900, 9},
		{990, 99},
	}
	for _, test := range tests {
		x, want := big.NewInt(test.fwd), big.NewInt(test.rev)
		if got := Reverse(x); got.Cmp(want) != 0 {
			t.Errorf("Reverse(%d): got %d, want %d", x, got, want)
		}
	}
}
