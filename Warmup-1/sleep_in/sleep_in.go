package main

import (
	"fmt"
	"os"
	"strconv"
)

// The parameter weekday is True if it is a weekday, and the parameter
// vacation is True if we are on vacation. We sleep in if it is not a
// weekday or we're on vacation. Return True if we sleep in.
func sleep_in(weekday, vacation bool) bool {
	return (vacation || !weekday)
}

func main() {
	weekday, _ := strconv.ParseBool(os.Args[1])
	vacation, _ := strconv.ParseBool(os.Args[2])
	fmt.Println(sleep_in(weekday, vacation))
}
