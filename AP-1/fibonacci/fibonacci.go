package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i < 8; i++ {
		fmt.Printf("%v) %v \n", i, fibonacci(i))
	}

}
