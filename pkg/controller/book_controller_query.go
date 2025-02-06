package controller

import (
	_const "GoProjects/pkg/const"
	"GoProjects/pkg/functions_to_import"
	_interface "GoProjects/pkg/interface"
	"encoding/json"
	"log/slog"
	"net/http"
)

type QueryController struct {
	repo _interface.IBookRepository
}

func NewQ(repo _interface.IBookRepository) *QueryController {
	slog.Info("Создаю новый бук квери контроллер")
	slog.Info("Создал бук квери контроллер")
	return &QueryController{
		repo: repo,
	}
}

// GetBooks godoc
//
// @Summary      Получить все книги
// @Description  Получить
// @Tags         books
// @Produce      json
// @Success      200  {object}  []models.Book
// @Failure      500  {string}  string  "Error while trying to find books or Failed to encode book data"
// @Router       /books [get]
func (c *QueryController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.repo.GetAllBooks()

	if err != nil {
		http.Error(w, "Error while trying to find books", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode book data", http.StatusInternalServerError)
	}
}

// GetBookById godoc
//
// @Summary      Получить книгу по айди
// @Description  Получить книгу по айди
// @Tags         books
// @Produce      json
// @Param        Book_id   path      int  true  "Book id"
// @Success      200  {object}  models.Book
// @Failure      400  {string}  string  "Missing book ID"
// @Failure      500  {string}  string  "Error while trying to find book or Failed to encode book data"
// @Router       /book/{id} [get]
func (c *QueryController) GetBookById(w http.ResponseWriter, r *http.Request) {

	id, _, constErr := functions_to_import.GetId(r)

	if constErr != _const.No_Error {
		return
	}

	book, err := c.repo.GetBookById(id)
	if err != nil {
		http.Error(w, "Error while trying to find book", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Failed to encode book data", http.StatusInternalServerError)
	}
}
