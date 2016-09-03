package main

import (
	"fmt"
)

func main() {
	var kk int
	kk = 222

	if kk, err := getMy(); err != nil {
		fmt.Printf("Error, %d,", kk)
	} else {
		fmt.Printf("Ok, %d,", kk)
	}
	fmt.Printf("After, %d,", kk)
}

func getMy() (int, error) {
	return 1111, nil
}
