package repository

import (
	"context"

	"github.com/bootcamp-go/go-web/internal"
)

// NewRepositoryTicketMock creates a new repository for tickets in a map
func NewRepositoryTicketMock() *RepositoryTicketMock {
	return &RepositoryTicketMock{}
}

// RepositoryTicketMock implements the repository interface for tickets
type RepositoryTicketMock struct {
	// FuncGetAll represents the mock for the GetTotalTickets function
	FuncGetAll func() (map[int]internal.TicketAttributes, error)

	FuncGetTotalTickets func() (i int, err error)
	// FuncGetTicketsByDestinationCountry
	FuncGetTicketsByDestinationCountry func(country string) (t map[int]internal.TicketAttributes, err error)

	// Spy verifies if the methods were called
	Spy struct {
		// Get represents the spy for the Get function
		Get int
		// GetTicketsByDestinationCountry represents the spy for the GetTicketsByDestinationCountry function
		GetTicketsByDestinationCountry int
	}
}

// GetAll returns all the tickets
func (r *RepositoryTicketMock) GetAll() (map[int]internal.TicketAttributes, error) {
	// spy
	r.Spy.Get++

	// mock
	i, err := r.FuncGetAll()
	return i, err
}

// GetTotalTickets returns all the tickets
func (r *RepositoryTicketMock) GetTotalTickets() (int, error) {
	i, err := r.FuncGetTotalTickets()
	return i, err
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMock) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	// spy
	r.Spy.GetTicketsByDestinationCountry++

	// mock
	t, err = r.FuncGetTicketsByDestinationCountry(country)
	return
}
