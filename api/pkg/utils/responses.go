package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func JSON(w http.ResponseWriter, statusCode int, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(val); err != nil {
		fmt.Printf("%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, errorResponse{
			Error: err.Error(),
			Code:  statusCode,
		})
	}
	JSON(w, http.StatusBadRequest, nil)
}
