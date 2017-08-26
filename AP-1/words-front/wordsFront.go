package main

import "fmt"

// Given an array of strings, return a new array
// containing the first N strings. N will be in the range 1..length.
func wordsFront(a []string, n int) []string {
	return a[:n]
}

func main() {
	//tests
	fmt.Println(wordsFront([]string{"a", "b", "c", "d"}, 1)) //["a"]
	fmt.Println(wordsFront([]string{"a", "b", "c", "d"}, 2)) //["a","b"]
	fmt.Println(wordsFront([]string{"a", "b", "c", "d"}, 3)) //["a","b","c"]
}
