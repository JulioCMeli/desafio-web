package repository

import "github.com/bootcamp-go/go-web/internal"

type TicketRepository interface {
	GetAll() (map[int]internal.TicketAttributes, error)
	GetTotalTickets() (int, error)
	GetTicketsByDestinationCountry(country string) (t map[int]internal.TicketAttributes, err error)
}
