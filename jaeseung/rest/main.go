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
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	portNum  *int
	repoType *string
	logLevel *string
)

func init() {
	portNum = flag.Int("port", 8888, "Server port")
	repoType = flag.String("repoType", "memory", "DB type(memory, mariaDB)")
	logLevel = flag.String("logLevel", "debug", "DB type(debug, info, warn)")
	flag.Parse()
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	switch *logLevel {
	case "info":
		log.Warn("log level=InfoLevel")
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.Warn("log level=WarnLevel")
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.Warn("log level=DebugLevel")
		log.SetLevel(log.DebugLevel)
	default:
		log.Warn("log level set to default=DebugLevel")
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	mux := mux.NewRouter()
	// / Home
	h := home.New()
	h.SetupRoutes(mux)

	// /vin
	repo, err := loadRepository(*repoType)
	if err != nil {
		log.Fatal(err)
	}
	v := vin.New(repo)
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
