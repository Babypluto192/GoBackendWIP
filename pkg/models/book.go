package models

// Book model info
// @Description Book information
// @Description with book id and name description and author
type Book struct {
	BookId      int
	Name        string
	Description string
	Author      string
}

// AddBook  model info
// @Description same as book but without id
type AddBook struct {
	Name        string
	Description string
	Author      string
}
