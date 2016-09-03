package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	s = "0.1"

}

func isNumber(s string) bool {
	arr := strings.Split(s, "e")
	if len(arr) > 2 {
		return false
	}

	if len(arr) == 2 {
		if !isArab(arr[1]) {
			return false
		}
		s = arr[0]
	}

	arr := strings.Split(s, ".")
	if len(arr) > 2 {
		return false
	}

	if len(arr) == 2 {
		if !isArab(arr[1]) {
			return false
		}
		s = arr[0]
		if len(s) == 0 {
			return true
		}
	} else {
		if len(s) == 0 {
			return false
		}
	}

	if !isArab(s) {
		return false
	}
	return true
}

func isArab(s1 string) bool {
	matched, err := regexp.MatchString("[0-9]*", s1)
	if err != nil {
		return false
	}
	if !matched {
		return false
	}
	return true
}
