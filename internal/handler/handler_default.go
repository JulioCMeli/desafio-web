package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/go-web/internal/service"
	"github.com/go-chi/chi/v5"
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

func (h *TicketHandler) GetTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dest := chi.URLParam(r, "dest")

		data, err := h.sv.GetTicketsByDestinationCountry(dest)
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

func (h *TicketHandler) GetAverageByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dest := chi.URLParam(r, "dest")

		data, err := h.sv.GetAverageByDestinationCountry(dest)
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
