package main

import (
	"reflect"
	"testing"
)

func TestFamilies(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{
			in:   "51404",
			want: []string{"514*4", "5*404", "51*04", "51*0*", "5140*", "*1404"},
		},
	}
	for _, test := range tests {
		got := families(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("families(%q) = %v, want %v", test.in, got, test.want)
		}
	}
}
