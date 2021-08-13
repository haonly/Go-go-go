package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Config interface {
	ServerAddress() string
	ServerPort() int
}

func NewServer(c Config, r *mux.Router) http.Server {
	srv := http.Server{
		Addr:	fmt.Sprintf("%s:%d", c.ServerAddress(), c.ServerPort()),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: r,
	}
	return srv
}