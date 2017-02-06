package main

import "fmt"

type agonal struct {
	value      int
	delta      int
	deltadelta int
}

func newAgonal(sides int) *agonal {
	return &agonal{delta: 1, deltadelta: sides - 2}
}

func (a *agonal) next() {
	a.value += a.delta
	a.delta += a.deltadelta
}

func main() {
	t := newAgonal(3)
	p := newAgonal(5)
	h := newAgonal(6)
	for ; ; h.next() {
		v := h.value
		for p.value < v {
			p.next()
		}
		for t.value < v {
			t.next()
		}
		if v == p.value && v == t.value {
			fmt.Println(v)
			if v == 0 || v == 1 || v == 40755 {
				continue
			}
			return
		}
	}
}
