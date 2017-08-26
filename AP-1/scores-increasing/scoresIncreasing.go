package main

import "fmt"

// Given an array of scores, return true if each
// score is equal or greater than the one before.
// The array will be length 2 or more.
func scoreIncreasing(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func main() {
	//tests
	fmt.Println(scoreIncreasing([]int{1, 3, 4})) //true
	fmt.Println(scoreIncreasing([]int{1, 3, 2})) //false
	fmt.Println(scoreIncreasing([]int{1, 1, 4})) //true
}
