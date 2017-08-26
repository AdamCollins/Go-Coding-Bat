package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzBuzz(n int) string {
	str := ""
	//Multiple of 3
	if n%3 == 0 {
		str += "Fizz"
	}
	//Multiple of 5
	if n%5 == 0 {
		str += "Buzz"
	}
	//Not a multiple
	if len(str) <= 0 {
		str = strconv.Itoa(n)
	}
	return str
}

func main() {
	sta, _ := strconv.ParseInt(os.Args[1], 10, 64)
	end, _ := strconv.ParseInt(os.Args[2], 10, 64)
	for i := int(sta); i <= int(end); i++ {
		fmt.Println(fizzBuzz(i))
	}
}
