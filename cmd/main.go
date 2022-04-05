package main

import (
	"fmt"

	"github.com/shuhei0818/algorithm"
)

func main() {
	q := algorithm.NewQueue[string]()

	err := q.Enqueue("a")
	if err != nil {
		panic(err)
	}

	err = q.Enqueue("b")
	if err != nil {
		panic(err)
	}

	err = q.Enqueue("c")
	if err != nil {
		panic(err)
	}

	d, err := q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	d, err = q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	d, err = q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	d, err = q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	d, err = q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

}
