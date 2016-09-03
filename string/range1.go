package main

import (
	"fmt"
	"time"
)

func main() {
	for k, v := range getSet() {
		fmt.Printf("oo k %v, v %v.\n", k, v)
		time.Sleep(time.Second * 2)
	}
}

func getSet() []int {
	fmt.Println("oo")
	return []int{11111123, 12222223, 12333333}
}
