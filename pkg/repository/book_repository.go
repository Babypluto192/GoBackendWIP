package repository

import (
	"GoProjects/pkg/db"
	"GoProjects/pkg/models"
	"context"
	"errors"
	"log/slog"
)

type BookRepository struct {
	db *db.DB
}

func New(db *db.DB) *BookRepository {
	slog.Debug("Создаю репоситорий с дб", db)
	slog.Info("Создаю репоситорий")
	slog.Info("Репоситорий создан")
	return &BookRepository{db}
}

func (repo *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	rows, err := repo.db.Pool.Query(
		context.Background(),
		"SELECT book_id, name, description, author FROM book",
	)
	if err != nil {
		return books, err
	}
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BookId,
			&book.Name,
			&book.Description,
			&book.Author); err != nil {
		}
		books = append(books, book)
	}
	return books, nil
}

func (repo *BookRepository) GetBookById(id string) (models.Book, error) {
	var book models.Book

	row := repo.db.Pool.QueryRow(
		context.Background(),
		"SELECT book_id, name, description, author FROM book WHERE book_id = $1",
		id,
	)

	err := row.Scan(
		&book.BookId,
		&book.Name,
		&book.Description,
		&book.Author,
	)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (repo *BookRepository) CreateBook(book models.AddBook) (bool, error) {
	_, err := repo.db.Pool.Exec(context.Background(), "INSERT INTO book(name, description, author) VALUES($1,$2,$3)", book.Name, book.Description, book.Author)
	if err != nil {
		return false, err
	}
	return true, nil

}
func (repo *BookRepository) UpdateBook(id string, book models.AddBook) (bool, error) {
	row := repo.db.Pool.QueryRow(context.Background(), "UPDATE book SET name = $2,description = $3, author = $4 WHERE book_id = $1   RETURNING name,  description, author", id, book.Name, book.Description, book.Author)

	var returnBook models.AddBook
	err := row.Scan(
		&returnBook.Name,
		&returnBook.Description,
		&returnBook.Author,
	)
	if err != nil {
		return false, err
	}

	if returnBook.Name != "" && returnBook.Description != "" && returnBook.Author != "" {
		return true, nil
	}

	return false, errors.New("no book updated")
}
func (repo *BookRepository) DeleteBook(id string) (bool, error) {
	_, err := repo.db.Pool.Exec(context.Background(), "DELETE FROM book where book_id = $1", id)
	if err != nil {
		return false, err
	}
	return true, nil
}
