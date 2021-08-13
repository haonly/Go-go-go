package main

import (
	"Go-go-go/hayeon-kim/project-registration/apis/registration"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

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

	router.HandleFunc("/api/v1/register", registration.Register).Methods("POST")
	router.HandleFunc("/api/v1/search", registration.Search).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
