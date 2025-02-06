package service

import (
	"GoProjects/pkg/const"
	"GoProjects/pkg/functions_to_import"
	_interface "GoProjects/pkg/interface"
	"log/slog"
	"net/http"
)

type Service struct {
	repo _interface.IBookRepository
}

func New(repo _interface.IBookRepository) *Service {
	slog.Info("Создаю Бук сервис")
	slog.Info("Бук сервис создан")
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddBook(request *http.Request) (bool, error, _const.Errors) {

	book, er, ex := functions_to_import.ParseBody(request)

	if ex != _const.No_Error {
		return false, er, _const.Decode_Error
	}

	isAdded, err := s.repo.CreateBook(book)

	if !isAdded {
		return false, nil, _const.Bad_Request
	}

	if err != nil {
		return false, err, _const.Server_Error
	}

	return true, nil, _const.No_Error
}

func (s *Service) UpdateBook(request *http.Request) (bool, error, _const.Errors) {
	book, er, ex := functions_to_import.ParseBody(request)

	if ex != _const.No_Error {
		return false, er, _const.Decode_Error
	}

	id, _, constErr := functions_to_import.GetId(request)

	if constErr != _const.No_Error {
		return false, error(nil), constErr
	}

	isUpdated, er := s.repo.UpdateBook(id, book)

	if !isUpdated {
		return false, nil, _const.Bad_Request
	}

	if er != nil {
		return false, er, _const.Server_Error
	}

	return true, nil, _const.No_Error
}

func (s *Service) DeleteBook(request *http.Request) (bool, error, _const.Errors) {

	id, _, constErr := functions_to_import.GetId(request)

	if constErr != _const.No_Error {
		return false, error(nil), constErr
	}

	isDeleted, err := s.repo.DeleteBook(id)

	if err != nil {
		return false, err, _const.Server_Error
	}

	if !isDeleted {
		return false, err, _const.Bad_Request
	}

	return true, nil, _const.No_Error
}
