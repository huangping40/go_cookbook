package main

import (
	"fmt"
	"github.com/foize/go.sgr"
)

func main() {
	sgr.Println("This is an example: [fg-red bold] important text [reset] normal text again.")

	// use sgr.Printf like you're used to with fmt
	// note: you should use `[reset]`  before the newline `\n`
	sgr.Printf("The secret number is [fg-17]%d [reset]\n", 42)

	// you can also parse once and re-use the parsed format string
	// using Parseln eliminates the need for a `[reset]\n` at the end of line
	secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
	fmt.Printf(secretNumberFormat, 42)

	fmt.Println("This text is normal again, MustParseln puts a reset at the end of the line.")
}
