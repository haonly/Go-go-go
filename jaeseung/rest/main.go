package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/rest/home"
	"github.com/haonly/Go-go-go/jaeseung/rest/server"
	vin "github.com/haonly/Go-go-go/jaeseung/rest/vin"
	"log"
	"os"
	"strconv"
)

var (
	portNum *int
)

func init() {
	portNum = flag.Int("port", 8888, "Server port")
	flag.Parse()
}

func main() {
	logger := log.New(os.Stdout, "[motor] ", log.LstdFlags|log.Lshortfile)
	mux := mux.NewRouter()
	h := home.New(logger)
	h.SetupRoutes(mux)
	v := vin.New(logger, vin.NewMemRepository())
	v.SetupRoutes(mux)

	srv := server.New(mux, ":"+strconv.Itoa(*portNum))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start %v", err)
	}
}
