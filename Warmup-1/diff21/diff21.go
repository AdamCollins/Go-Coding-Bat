package main

import (
	"fmt"
	"os"
	"strconv"
)

// Given an int n, return the absolute difference
// between n and 21, except return double the absolute
// difference if n is over 21.
func diff21(n int64) int64 {
	return 21 - n
}

func main() {
	n, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Println(diff21(n))
}
