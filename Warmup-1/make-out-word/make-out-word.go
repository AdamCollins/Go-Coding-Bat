package main

import (
	"fmt"
	"os"
)

// Given an "out" string length 4, such as "<<>>",
// and a word, return a new string where the word
// is in the middle of the out string, e.g. "<<word>>".
func makeOutWord(out, word string) string {
	sta := out[0:2]
	end := out[2:4]
	return sta + word + end
}

func main() {
	fmt.Println(makeOutWord(os.Args[1], os.Args[2]))
}
