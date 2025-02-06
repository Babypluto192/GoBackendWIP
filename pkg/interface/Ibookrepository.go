package _interface

import "GoProjects/pkg/models"

type IBookRepository interface {
	GetAllBooks() ([]models.Book, error)

	GetBookById(id string) (models.Book, error)

	CreateBook(book models.AddBook) (bool, error)

	UpdateBook(id string, book models.AddBook) (bool, error)

	DeleteBook(id string) (bool, error)
}
