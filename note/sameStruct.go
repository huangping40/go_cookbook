package main

import (
	"fmt"

	"api"
)

func main() {
	var a struct {
		x int    `x`
		s string `s`
	}

	var b struct {
		x int
		s string
	}

	c := api.a
	fmt.Printf("%+v", a)
	fmt.Printf("%+v", c)
	//b = a

	fmt.Printf("%+v", b)
}
