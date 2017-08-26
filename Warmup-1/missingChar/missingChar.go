package main

import (
	"fmt"
	"os"
	"strconv"
)

// Given a non-empty string and an int n,
// return a new string where the char at index
// n has been removed. The value of n will
// be a valid index of a char in the original
// string (i.e. n will be in the
// range 0..len(str)-1 inclusive).
func missingChar(str string, n int64) string {
	front := str[0:n]
	back := str[n+1 : len(str)]
	return front + back
}

func main() {
	str := os.Args[1]
	n, _ := strconv.ParseInt(os.Args[2], 10, 64)
	fmt.Println(missingChar(str, n))
}
