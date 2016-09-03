package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	k := hashCode("abcdredhat_234")
	fmt.Printf("K: %v\n", k)

	k = hashCode("abcchch吃饭dredhat_234")
	fmt.Printf("K: %v\n", k)

	p := seededHash(k, "12na吃饭")
	fmt.Printf("p: %v\n", p)

	p = seededHash(k, "1223455688")
	fmt.Printf("p: %v\n", p)

}

func hashCode(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func seededHash(seed uint32, s string) uint32 {
	var h uint32 = seed % 100
	for i := 0; i < len(s); i++ {
		h = ((h * 31) + uint32(s[i])) % 100
	}
	return h
}
