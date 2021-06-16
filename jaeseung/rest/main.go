package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/rest/home"
	"github.com/haonly/Go-go-go/jaeseung/rest/server"
	vin "github.com/haonly/Go-go-go/jaeseung/rest/vin"
	"github.com/haonly/Go-go-go/jaeseung/rest/vin/repo/mariadb"
	"github.com/haonly/Go-go-go/jaeseung/rest/vin/repo/memory"
	"log"
	"os"
	"strconv"
)

var (
	portNum  *int
	repoType *string
)

func init() {
	portNum = flag.Int("port", 8888, "Server port")
	repoType = flag.String("repoType", "memory", "DB type(memory, mariaDB)")
	flag.Parse()
}

func main() {
	logger := log.New(os.Stdout, "[motor] ", log.LstdFlags|log.Lshortfile)
	mux := mux.NewRouter()
	// / Home
	h := home.New(logger)
	h.SetupRoutes(mux)

	// /vin
	repo, err := loadRepository(*repoType)
	if err != nil {
		log.Fatal(err)
	}
	v := vin.New(logger, repo)
	v.SetupRoutes(mux)

	srv := server.New(mux, ":"+strconv.Itoa(*portNum))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start %v", err)
	}
}

func loadRepository(repoType string) (vin.VinRepository, error) {
	if repoType == "memory" {
		return memory.New(), nil
	}
	if repoType == "mariaDB" {
		return mariadb.New(), nil
	}
	return nil, fmt.Errorf("Unknown repository type repoType=%s", repoType)
}
