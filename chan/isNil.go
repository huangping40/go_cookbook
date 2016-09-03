package main

import (
	"fmt"
)

type exampleMapper struct {
	keys [][]byte
	vals [][]byte
	id   string
}

func main() {
	vals := []exampleMapper{}

	fmt.Printf("kk: %v \n", len(vals))
}
