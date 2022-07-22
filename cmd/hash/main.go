package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	h := fnv.New32()
	h.Write([]byte("hello world.....aaaaあwwwsawasefrew"))
	sum := h.Sum32()

	fmt.Printf("sum: %v\n", sum)
}
