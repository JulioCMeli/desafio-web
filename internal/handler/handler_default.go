package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/go-web/internal/service"
)

// TicketHandler is a handler with a map of users as data
type TicketHandler struct {
	sv service.ServiceTicketDefault
}

// MyResponse is an struct for the response
type MyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewHandler returns a new TicketHandler
func NewTicketHandler(sv *service.ServiceTicketDefault) *TicketHandler {

	var myHandler TicketHandler
	myHandler.sv = *sv
	return &myHandler
}

// Get returns a handler for the GET /products route
func (h *TicketHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := h.sv.GetAll()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		code := http.StatusOK
		body := MyResponse{Message: "OK", Data: data}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}
