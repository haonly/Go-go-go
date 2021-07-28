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
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/module/apmlogrus"
	"go.elastic.co/ecslogrus"
	"os"
)

var (
	logLevel *string
)

func init() {
	logLevel = flag.String("logLevel", "debug", "DB type(debug, info, warn)")
	flag.Parse()
	log.SetFormatter(&ecslogrus.Formatter{
		DataKey: "labels", // For ECS compliance
	})
	log.SetReportCaller(true)
	log.AddHook(&apmlogrus.Hook{})

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
	os.Setenv("ELASTIC_APM_SERVER_URL", "0.0.0.0:8201")
	os.Setenv("RepositoryType", "mariaDB")
}

func main() {
	cfg := config.New()

	repo, err := loadRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	mux := mux.NewRouter()
	apmgorilla.Instrument(mux)
	h := home.New()
	h.SetupRoutes(mux)

	v := vin.New(repo)
	v.SetupRoutes(mux)

	srv := server.New(cfg, mux)
	log.Info("Before Server has started")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start %v", err)
	}
	log.Info("Server has started")
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
