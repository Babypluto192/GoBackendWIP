package tests

import (
	controller2 "GoProjects/pkg/controller"
	"GoProjects/pkg/interface/mocks"
	"GoProjects/pkg/models"
	"GoProjects/pkg/service"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBooksShouldReturnOk(t *testing.T) {

	mockRepo := new(mocks.IBookRepository)
	controller := controller2.NewQ(mockRepo)

	expectedBooks := make([]models.Book, 2)
	mockRepo.On("GetAllBooks").Return(expectedBooks, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	controller.GetBooks(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusOK, response.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetAllBooksShouldReturnStatus500WhenError(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	controller := controller2.NewQ(mockRepo)

	mockRepo.On("GetAllBooks").Return([]models.Book{}, errors.New("some error"))

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	controller.GetBooks(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetBookByIdShouldReturnOk(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	controller := controller2.NewQ(mockRepo)

	mockRepo.On("GetBookById", mock.Anything).Return(models.Book{}, nil)

	request := httptest.NewRequest(http.MethodGet, "/book/1", nil)
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.GetBookById)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusOK, response.Code)
	mockRepo.AssertExpectations(t)

}

func TestGetBookByIdShouldReturnStatus500WhenError(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	controller := controller2.NewQ(mockRepo)

	mockRepo.On("GetBookById", mock.Anything).Return(models.Book{}, errors.New("some error"))

	request := httptest.NewRequest(http.MethodGet, "/book/1", nil)
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.GetBookById)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockRepo.AssertExpectations(t)
}
func TestCreateBookShouldReturnOk(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	addBook := models.AddBook{}
	mockRepo.On("CreateBook", mock.Anything).Return(true, nil)

	requestBody, err := json.Marshal(addBook)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	controller.CreateBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusCreated, response.Code)
	mockRepo.AssertExpectations(t)

}

func TestCreateBookShouldReturnStatus500WhenError(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	requestBody, err := json.Marshal(models.AddBook{
		Name:        "",
		Description: "",
		Author:      "",
	})
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	mockRepo.On("CreateBook", mock.Anything).Return(true, errors.New("some error"))
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	controller.CreateBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockRepo.AssertExpectations(t)

}

func TestCreateBookShouldReturnStatus400WhenBookIsntCreated(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	requestBody, err := json.Marshal(models.AddBook{
		Name:        "",
		Description: "",
		Author:      "",
	})
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	mockRepo.On("CreateBook", mock.Anything).Return(false, errors.New("some error"))
	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	controller.CreateBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateBookShouldReturnOk(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	addBook := models.AddBook{}
	mockRepo.On("UpdateBook", mock.Anything, mock.Anything).Return(true, nil)

	requestBody, err := json.Marshal(addBook)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	request := httptest.NewRequest(http.MethodPut, "/book/1", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.UpdateBook)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	controller.UpdateBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusOK, response.Code)
	mockRepo.AssertExpectations(t)

}

func TestUpdateBookShouldReturnStatus500WhenBookIsntUpdated(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	addBook := models.AddBook{}
	mockRepo.On("UpdateBook", mock.Anything, mock.Anything).Return(true, errors.New("some error"))
	requestBody, err := json.Marshal(addBook)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	request := httptest.NewRequest(http.MethodPut, "/book/1", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.UpdateBook)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	controller.UpdateBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockRepo.AssertExpectations(t)
}
func TestUpdateBookShouldReturnStatus400WhenBookIsntUpdated(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)
	addBook := models.AddBook{}
	mockRepo.On("UpdateBook", mock.Anything, mock.Anything).Return(false, nil)
	requestBody, err := json.Marshal(addBook)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	request := httptest.NewRequest(http.MethodPut, "/book/1", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.UpdateBook)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	controller.UpdateBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBookShouldReturnNoContent(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	mockRepo.On("DeleteBook", mock.Anything).Return(true, nil)

	request := httptest.NewRequest(http.MethodDelete, "/book/1", nil)
	request.Header.Set("Content-Type", "application/json")
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.DeleteBook)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	controller.DeleteBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusNoContent, response.Code)
	mockRepo.AssertExpectations(t)

}

func TestDeleteBookShouldReturnStatus500WhenError(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	mockRepo.On("DeleteBook", mock.Anything).Return(true, errors.New("some error"))
	request := httptest.NewRequest(http.MethodDelete, "/book/1", nil)
	request.Header.Set("Content-Type", "application/json")
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.DeleteBook)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	controller.DeleteBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBookShouldReturnStatus400WhenBookIsntDeleted(t *testing.T) {
	mockRepo := new(mocks.IBookRepository)
	service2 := service.New(mockRepo)
	controller := controller2.New(service2)

	mockRepo.On("DeleteBook", mock.Anything).Return(false, nil)

	request := httptest.NewRequest(http.MethodDelete, "/book/1", nil)
	request.Header.Set("Content-Type", "application/json")
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.DeleteBook)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	controller.DeleteBook(response, request)
	result := response.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockRepo.AssertExpectations(t)
}
