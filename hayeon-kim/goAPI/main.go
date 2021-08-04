package main

import (
	"Go-go-go/hayeon-kim/goAPI/apis/vehicle"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/vehicle/all/findall", vehicle.FindAll).Methods("GET")
	router.HandleFunc("/api/vehicle/{keyword}", vehicle.Search).Methods("GET")
	router.HandleFunc("/api/vehicle/vin", vehicle.Create).Methods("POST")
	router.HandleFunc("/api/vehicle/vin", vehicle.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
