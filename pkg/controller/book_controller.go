package controller

import (
	_const "GoProjects/pkg/const"
	"GoProjects/pkg/service"
	"log/slog"
	"net/http"
)

type Controller struct {
	ser *service.Service
}

func New(ser *service.Service) *Controller {
	slog.Info("Создаю новый бук сервис")
	slog.Info("Создал")
	return &Controller{
		ser: ser,
	}
}

// CreateBook godoc
//
// @Summary Добавить книгу
// @Description Добавить книгу в БД
// @Tags books
// @accept json
// @Param  book body models.AddBook true "Add book"
// @Success 201 {string} string "Book created successfully"
// @Failure 500 {string} string "Error data"
// @Failure 400 {string} string "Failed to decode book data"
// @Router  /books [post]
func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	isAdded, err, errorType := c.ser.AddBook(r)

	if errorType == _const.Decode_Error || errorType == _const.Server_Error {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorType == _const.Bad_Request || !isAdded {
		http.Error(w, "Book was Not Added", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte("Book created successfully"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// UpdateBook godoc
//
// @Summary Обновить книгу
// @Description Обновляет книгу по айди
// @Tags books
// @accept json
// @Produce json
// @Param book body models.AddBook true "Update book"
// @Param id path int true "Book id"
// @Success 200 {object} models.AddBook
// @Failure 400 {string} string "Failed to decode book data or Missing book ID"
// @Failure 500 {string} string "Failed to encode book data or error message"
// @Router /api/books/{id} [put]
func (c *Controller) UpdateBook(w http.ResponseWriter, r *http.Request) {

	isUpdated, err, errorType := c.ser.UpdateBook(r)

	if errorType == _const.Decode_Error || errorType == _const.Server_Error {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorType == _const.Bad_Request || !isUpdated {
		http.Error(w, "Book was not updated", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte("Book updated successfully"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// DeleteBook godoc
//
// @Summary Удалить книгу
// @Description Удаляет книгу по айди
// @Tags books
// @Param id path int true "Book id"
// @Success 204
// @Failure 400 {string} string "Error while deleting book or Missing book ID"
// @Failure 500 {string} string "error message"
// @Router /api/books/{id} [delete]
func (c *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	isDeleted, err, errorType := c.ser.DeleteBook(r)

	if errorType == _const.Decode_Error || errorType == _const.Server_Error {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorType == _const.Bad_Request || !isDeleted {
		http.Error(w, "Book was not deleted", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
