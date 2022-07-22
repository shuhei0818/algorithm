package main

import (
	"fmt"

	"github.com/shuhei0818/algorithm/list"
)

func main() {
	s, err := list.NewQueue[string](10)
	if err != nil {
		panic(err)
	}
	s.Add("a")
	s.Add("b")
	s.Add("c")

	d, err := s.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %v\n", d)

	d, err = s.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %v\n", d)

	d, err = s.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %v\n", d)
}
