package functions_to_import

import (
	_const "GoProjects/pkg/const"
	"GoProjects/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetId(r *http.Request) (string, error, _const.Errors) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		return "", error(nil), _const.Bad_Request
	}
	return id, error(nil), _const.No_Error
}

func ParseBody(r *http.Request) (models.AddBook, error, _const.Errors) {
	var book models.AddBook

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		return book, err, _const.Decode_Error
	}

	return book, nil, _const.No_Error
}
