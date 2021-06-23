package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/rest/config"
	"github.com/haonly/Go-go-go/jaeseung/rest/home"
	"github.com/haonly/Go-go-go/jaeseung/rest/server"
	vin "github.com/haonly/Go-go-go/jaeseung/rest/vin"
	"github.com/haonly/Go-go-go/jaeseung/rest/vin/repo/mariadb"
	"github.com/haonly/Go-go-go/jaeseung/rest/vin/repo/memory"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	logLevel *string
)

func init() {
	logLevel = flag.String("logLevel", "debug", "DB type(debug, info, warn)")
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

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
	cfg := config.New()

	repo, err := loadRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	mux := mux.NewRouter()
	h := home.New()
	h.SetupRoutes(mux)

	v := vin.New(repo)
	v.SetupRoutes(mux)

	srv := server.New(cfg, mux)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start %v", err)
	}
}

func loadRepository(c *config.Config) (vin.VinRepository, error) {
	if c.RepositoryType == "memory" {
		return memory.New(), nil
	}
	if c.RepositoryType == "mariaDB" {
		return mariadb.New(c), nil
	}
	return nil, fmt.Errorf("Unknown repository type repoType=%s", c.RepositoryType)
}
