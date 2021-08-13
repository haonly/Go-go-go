package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/config"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/controller"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/entity/repo"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/entity/repo/memory"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/usecase"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/server"
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

	uc := usecase.NewReview(repo)
	ctrl := controller.NewReview(uc)
	ctrl.SetupRoutes(mux)

	srv := server.New(cfg, mux)
	log.Info("Before Server has started")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server can't start %v", err)
	}
	log.Info("Server has started")
}

func loadRepository(c *config.Config) (repo.ReviewRepository, error) {
	if c.RepositoryType == "memory" {
		return memory.New(), nil
	}
	return nil, fmt.Errorf("Unknown repository type repoType=%s", c.RepositoryType)
}
