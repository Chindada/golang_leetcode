package main

import (
	"fmt"
	"strconv"
)

func main() {
	// a := 1
	// b := 0
	// fmt.Println(a / b)

	_, err := strconv.ParseFloat("123.33", 64)
	fmt.Println(fmt.Errorf("XXX: %w", err))
	if fmt.Errorf("XXX: %w", err) == nil {
		fmt.Println("123")
	}
}
