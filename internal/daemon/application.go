package daemon

import (
	"net/http"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	handlers "github.com/japhy-tech/backend-test/internal/inbound/handlers/breeds"
	"github.com/japhy-tech/backend-test/internal/ports"
)

type App struct {
	logger *charmLog.Logger
	crud   ports.BreedsCRUD
}

func NewApp(logger *charmLog.Logger, crud ports.BreedsCRUD) *App {
	return &App{
		logger: logger,
		crud:   crud,
	}
}

func (a *App) RegisterRoutes(r *mux.Router, logger *charmLog.Logger) {
	r.HandleFunc("/", handlers.Handle(a.crud.Create, logger)).Methods(http.MethodPost)
	r.HandleFunc("/", handlers.HandleGet(a.crud, logger)).Methods(http.MethodGet).Queries("name", "{name}")
	r.HandleFunc("/", handlers.HandleGet(a.crud, logger)).Methods(http.MethodGet).Queries("id", "{id}")
	r.HandleFunc("/", handlers.HandleGet(a.crud, logger)).Methods(http.MethodGet).Queries("species", "{species}")
	r.HandleFunc("/", handlers.HandleGet(a.crud, logger)).Methods(http.MethodGet).Queries("pet_size", "{pet_size}")
	r.HandleFunc("/", handlers.HandleGet(a.crud, logger)).Methods(http.MethodGet).Queries("weight", "{weight}")
	r.HandleFunc("/", handlers.Handle(a.crud.Update, logger)).Methods(http.MethodPut)
	r.HandleFunc("/", handlers.HandleDelete(a.crud, logger)).Methods(http.MethodDelete).Queries("id", "{id}")
}
