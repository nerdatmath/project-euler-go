package main

import "fmt"

var primes [4096]int

func init() {
    wheel := [...]int{7, 11, 13, 17, 19, 23, 29, 31}
    for p, offset := primes[:], 0; len(p) >= len(wheel); p, offset = p[len(wheel):], offset+30 {
        for j, n := range wheel {
            p[j] = offset+n
        }
    }
    for i, p := range primes {
        for _, d := range primes[:i] {
            if d == 0 {
                continue
            }
            if p % d == 0 {
                primes[i] = 0
                break
            }
            if d * d > p {
                break
            }
        }
    }
}

func goldbach(x int) (int, int, bool) {
    for sq, d, dd := 0, 2, 4; sq < x; sq, d = sq+d, d+dd {
        for _, p := range primes {
            if p == 0 {
                continue
            }
            if sq+p == x {
                return sq, p, true
            }
        }
    }
    return 0, 0, false
}

func main() {
    for i := 9; ; i += 2 {
        sq, p, ok := goldbach(i)
        if !ok {
            fmt.Println()
            fmt.Println(i)
            return
        }
        fmt.Printf("\r%8d + %8d = %8d", sq, p, i)
    }
}
