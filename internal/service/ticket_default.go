package service

import (
	"github.com/bootcamp-go/go-web/internal"
	"github.com/bootcamp-go/go-web/internal/repository"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.TicketRepository
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.TicketRepository) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetAll returns the total tickets
func (s *ServiceTicketDefault) GetAll() (t map[int]internal.TicketAttributes, err error) {
	return s.rp.GetAll()
}

// GetTotalAmountTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	return s.rp.GetTotalTickets()
}
