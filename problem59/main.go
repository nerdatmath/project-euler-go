package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	ciphertext, err := ioutil.ReadFile("cipher.txt")
	if err != nil {
		panic(err)
	}
	numbers := strings.Split(string(ciphertext), ",")
	ciphertext = make([]byte, len(numbers))
	m := []map[string]int{make(map[string]int), make(map[string]int), make(map[string]int)}
	for i, n := range numbers {
		b, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		ciphertext[i] = byte(b)
		if c := byte(b) ^ ' '; 'a' <= c && c <= 'z' {
			m[i%3][string(c)] = m[i%3][string(c)] + 1
		}
	}
	key := make([]byte, 3)
	for i, mm := range m {
		ch := ""
		for c, cnt := range mm {
			if cnt > mm[ch] {
				ch = c
			}
		}
		key[i] = ch[0]
	}
	plaintext := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		plaintext[i] = b ^ key[i%3]
	}
	fmt.Println(string(plaintext))
}
