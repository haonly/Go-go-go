package app

import (
	"Go-go-go/hayeon-kim/registration/app/database"
	"github.com/gorilla/mux"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     database.RegistrationDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/registers", a.CreateRegistrationHandler()).Methods("POST")
}
