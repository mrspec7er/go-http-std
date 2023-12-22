package utils

import (
	"encoding/json"
	"net/http"
)

func InternalServerErrorHandler(w http.ResponseWriter, status int, err error) {
	message := err.Error()
    response := struct {
		Status		bool `json:"status"`
		Message 	*string `json:"message"`
		Data 		interface{} `json:"data"`
		Metadata 	interface{} `json:"metadata"`
	}{
		Status: false,
		Message: &message,
		Data: nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(status)
	w.Write(responseData)
}

func NotFoundHandler(w http.ResponseWriter) {
	message := "Data Not Found"
    response := struct {
		Status		bool `json:"status"`
		Message 	*string `json:"message"`
		Data 		interface{} `json:"data"`
		Metadata 	interface{} `json:"metadata"`
	}{
		Status: false,
		Message: &message,
		Data: nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(responseData)
}

func BadRequestHandler(w http.ResponseWriter) {
	message := "Bad Request"
    response := struct {
		Status		bool `json:"status"`
		Message 	*string `json:"message"`
		Data 		interface{} `json:"data"`
		Metadata 	interface{} `json:"metadata"`
	}{
		Status: false,
		Message: &message,
		Data: nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(responseData)
}

func UnauthorizeUser(w http.ResponseWriter) {
	message := "Unauthorize user"
    response := struct {
		Status		bool `json:"status"`
		Message 	*string `json:"message"`
		Data 		interface{} `json:"data"`
		Metadata 	interface{} `json:"metadata"`
	}{
		Status: false,
		Message: &message,
		Data: nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write(responseData)
}

func MutationSuccessResponse(w http.ResponseWriter, message string)  {
	response := struct {
		Status		bool `json:"status"`
		Message 	*string `json:"message"`
		Data 		interface{} `json:"data"`
		Metadata 	interface{} `json:"metadata"`
	}{
		Status: true,
		Message: &message,
		Data: nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func GetSuccessResponse(w http.ResponseWriter, message *string, data interface{}, metadata interface{})  {
	response := struct {
		Status		bool `json:"status"`
		Message 	*string `json:"message"`
		Data 		interface{} `json:"data"`
		Metadata 	interface{} `json:"metadata"`
	}{
		Status: true,
		Message: message,
		Data: data,
		Metadata: metadata,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, 500, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}