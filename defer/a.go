package main

import (
	"fmt"
	"time"
)

type status struct {
	ii *status2
}
type status2 struct {
	ii int32
}

func main() {

	defer get(time.Now())
	time.Sleep(5 * time.Second)
	fmt.Printf("ok in main: %v.\n", time.Now())
}

func get(begin time.Time) {
	fmt.Printf("ok in get: %v.\n", begin)
}
