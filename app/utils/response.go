package utils

import (
	"encoding/json"
	"net/http"
)

func InternalServerErrorHandler(w http.ResponseWriter, status int, err error) {
    w.WriteHeader(status)
    w.Write([]byte(err.Error()))
}

func NotFoundHandler(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("Not Found"))
}

func BadRequestHandler(w http.ResponseWriter) {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Bad Request"))
}

func UnauthorizeUser(w http.ResponseWriter) {
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte("Unauthorize User"))
}

func CreateSuccessResponse(w http.ResponseWriter, message string)  {
	responseJSON := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(responseJSON)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func GetSuccessResponse(w http.ResponseWriter, data interface{})  {
	responseJSON := data
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(responseJSON)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}