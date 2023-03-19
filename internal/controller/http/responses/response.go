package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error structure for ErrorResponse
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// Data structure for Response
type Data struct {
	Code int         `json:"code"`
	Args interface{} `json:"response"`
}

// Message structure for MessageResponse
type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response function return responses
func Response(w http.ResponseWriter, statusCode int, args interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := Data{
		Code: statusCode,
		Args: args,
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Print(err)
	}
}

// MessageResponse function returns an message and its status code
func MessageResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := Message{
		Code:    statusCode,
		Message: message,
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Print(err)
	}
}

// ErrorResponse function returns an error message and its status code
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := Error{
		Code:  statusCode,
		Error: message,
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Print(err)
	}
}
