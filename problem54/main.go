package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nerdatmath/project-euler-go/poker"
)

func main() {
	f, err := os.Open("poker.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open poker.txt: %v", err)
		os.Exit(1)
	}
	defer f.Close()
	out, err := os.Create("outcomes.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open outcomes.txt: %v", err)
		os.Exit(1)
	}
	defer out.Close()
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		hand1 := poker.ParseHand(line[0:14])
		hand2 := poker.ParseHand(line[15:29])
		eval := "<"
		if poker.Less(hand2.Evaluate(), hand1.Evaluate()) {
			eval = ">"
			count++
		}
		fmt.Fprintln(out, hand1, hand2, hand1.Evaluate(), eval, hand2.Evaluate())
	}
	fmt.Println(count)
}
