package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/bootcamp-go/go-web/internal"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
	lastId   int
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (mapi map[int]internal.TicketAttributes, lastId int) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		fmt.Println(err)
		return nil, 1
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	mapi = make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			fmt.Println(err)
			return mapi, 1
		}

		// serialize the record
		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   0,
		}

		if f, err := strconv.ParseFloat(record[5], 64); err == nil {
			ticket.Price = f
		}

		// add the ticket to the map
		idStr := record[0]
		id, err := strconv.Atoi(idStr)
		if err == nil {
			mapi[id] = ticket
			t.lastId = id
			//fmt.Println(t)
		}
	}

	return mapi, t.lastId
}
