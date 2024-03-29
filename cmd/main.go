package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bootcamp-go/go-web/internal/handler"
	"github.com/bootcamp-go/go-web/internal/loader"
	"github.com/bootcamp-go/go-web/internal/repository"
	"github.com/bootcamp-go/go-web/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	// env
	// ...
	os.Setenv("SERVER_ADDR", ":8080")
	os.Setenv("DB_FILE", "tickets.csv")

	// application
	// - config
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	l := loader.NewLoaderTicketCSV(a.dbFile)
	db, lastId := l.Load()
	rp := repository.NewRepositoryTicketMap(db, lastId)
	// service ...
	sv := service.NewServiceTicketDefault(rp)
	// handler ...
	h := handler.NewTicketHandler(sv)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	(*a).rt.Get("/ticket", h.GetAll())
	(*a).rt.Get("/ticket/getByCountry/{dest}", h.GetTicketsByDestinationCountry())
	(*a).rt.Get("/ticket/getAverage/{dest}", h.GetAverageByDestinationCountry())

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
