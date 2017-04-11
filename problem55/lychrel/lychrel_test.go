package lychrel_test

import (
	"testing"

	"github.com/nerdatmath/project-euler-go/problem55/lychrel"
)

func TestLychrel(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{196, true},
		{11, false},
		{349, false},
		{4994, true},
	}
	for _, test := range tests {
		got := lychrel.Lychrel(test.in)
		if got != test.want {
			t.Errorf("Lychrel(%d): got %t, want %t", test.in, got, test.want)
		}
	}
}
