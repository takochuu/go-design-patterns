package main

import (
	"fmt"

	"github.com/takochuu/go-design-patterns/composite"
	"github.com/takochuu/go-design-patterns/decorator"
	"github.com/takochuu/go-design-patterns/iterator"
)

func main() {
	// composite
	list := composite.Newlist()
	list.Add(composite.NewProductA())
	list.Add(composite.NewProductB())
	fmt.Println(list.GetAmount())

	// iterator
	bookshelf := iterator.BookShelf{}
	bookshelf.Add(&iterator.Book{Name: "Hoge"})
	bookshelf.Add(&iterator.Book{Name: "Fuga"})
	bookshelf.Add(&iterator.Book{Name: "Piyo"})
	i := bookshelf.Iterator()
	for i.HasNext() {
		book := i.Next()
		fmt.Println(book.Name)
	}

	// decorator
	cashewnuts := decorator.NewCashewNuts(decorator.NewMatcha())
	fmt.Println(cashewnuts.GetName())
	fmt.Println(cashewnuts.GetCost())
}
