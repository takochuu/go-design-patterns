package facade

var books = map[string]string{
	"hoge": "a",
	"fuga": "b",
}

type BookName struct {
}

func (b *BookName) getNameByTitle(title string) string {
	if _, ok := books[title]; !ok {
		return "not found"
	}
	return books[title]
}

type BookTitleMaker struct {
}

func NewBookTitleMaker() BookTitleMaker {
	return BookTitleMaker{}
}

func (b *BookTitleMaker) GetName(title string) string {
	bookName := BookName{}
	return bookName.getNameByTitle(title)
}
