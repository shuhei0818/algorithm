package main

import "fmt"

type Data struct {
	name string
}

type Cell struct {
	next *Cell
	data Data
}

func NewList(d Data) Cell {
	return Cell{
		next: nil,
		data: d,
	}
}

func (c *Cell) Add(d Data) {
	for c.next != nil {
		c = c.next
	}
	c.next = &Cell{next: nil, data: d}
}

func (c *Cell) Print() {
	for c.next != nil {
		fmt.Printf("c: %v\n", c)
		c = c.next
	}
	fmt.Printf("c: %v\n", c)
}

func main() {
	c := NewList(Data{name: "hello"})
	c.Add(Data{name: "world"})
	c.Add(Data{name: "world"})
	c.Add(Data{name: "world"})

	c.Print()
}
