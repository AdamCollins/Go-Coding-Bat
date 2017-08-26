package main

import (
	"fmt"
	"os"
)

// Given a string, return a new string where
// the first and last chars have been exchanged.
func front_back(str string) string {
	if len(str) <= 1 {
		return str
	}
	return string(str[len(str)-1]) + str[1:len(str)-1] + string(str[0])
}

func main() {
	fmt.Println(front_back(os.Args[1]))
}
