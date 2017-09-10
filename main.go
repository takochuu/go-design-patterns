package main

import (
	"fmt"
	"github.com/takochuu/go-design-patterns/composite"
)

func main() {
	list := composite.Newlist()
	list.Add(composite.NewProductA())
	list.Add(composite.NewProductB())
	fmt.Println(list.GetAmount())
}
