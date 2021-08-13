package main

import (
	"Go-go-go/hayeon-kim/project-registration/apis/registration"
	"Go-go-go/hayeon-kim/project-registration/config"
	"Go-go-go/hayeon-kim/project-registration/server"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()

	mux := mux.NewRouter()

	srv := server.NewServer(config, mux)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server not started %v", err)
	}
	log.Println("Server tarted!")

	/*
		{
			"company": "기아",
			"model": "MQ4",
			"type": "HEV",
			"year": 2021,
			"capacity": 3000
		}
	*/
	// 차 정보 json

	mux.HandleFunc("/api/v1/register", registration.Register).Methods("POST")
	mux.HandleFunc("/api/v1/search", registration.Search).Methods("GET")

	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
