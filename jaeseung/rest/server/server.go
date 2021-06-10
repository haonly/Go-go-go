package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func New(r *mux.Router, srvAddr string) http.Server {
	srv := http.Server{
		Addr:         srvAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	return srv
}
