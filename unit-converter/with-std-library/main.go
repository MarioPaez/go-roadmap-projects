package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"std-library/model"
	"std-library/service"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func main() {

	http.HandleFunc("/", conversionHandler)
	fmt.Println("server initialized in localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func conversionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	var conversion model.Conversion

	if err := json.NewDecoder(r.Body).Decode(&conversion); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}

	result, err := service.DoConversion(&conversion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}

	msg := fmt.Sprintf("Result of your calculation: %.3f %s = %.3f %s", conversion.Value, conversion.From, result, conversion.To)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: msg,
	})

}
