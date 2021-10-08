package web

import (
	"fmt"
	"net/http"
)

func WriteHttpError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	_, err := w.Write([]byte(http.StatusText(code)))
	if err != nil {
		LogError(err)
	}
}

func WriteNotFoundError(w http.ResponseWriter) {
	WriteHttpError(w, http.StatusNotFound)
}

func WriteBadRequestError(w http.ResponseWriter) {
	WriteHttpError(w, http.StatusBadRequest)
}

func WriteMethodNotAllowedError(w http.ResponseWriter) {
	WriteHttpError(w, http.StatusMethodNotAllowed)
}

func WriteServerError(w http.ResponseWriter, err error) {
	LogError(err)
	WriteHttpError(w, http.StatusInternalServerError)
}

func LogError(err error) {
	fmt.Printf("\nerror: %v\n", err)
}
