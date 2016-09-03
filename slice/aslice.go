package main

import (
	"fmt"
)

func main() {
	var a []int
	a = nil

	for _, c := range a {
		fmt.Printf("ff %v.", c)
	}
}
