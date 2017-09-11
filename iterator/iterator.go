package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() *Book
}

type BookShelf struct {
	Books []*Book
}

func (self *BookShelf) Add(book *Book) {
	self.Books = append(self.Books, book)
}

func (self *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: self}
}

type BookShelfIterator struct {
	BookShelf *BookShelf
	last      int
}

func (self *BookShelfIterator) HasNext() bool {
	if self.last >= len(self.BookShelf.Books) {
		return false
	}
	return true
}

func (self *BookShelfIterator) Next() *Book {
	book := self.BookShelf.Books[self.last]
	self.last++
	return book
}

type Book struct {
	Name string
}
