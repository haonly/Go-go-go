package api

import (
	"Go-go-go/ParkJiHyo/goAPI/api/service"
	"Go-go-go/ParkJiHyo/goAPI/store"
	"context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type ServerRun struct {
	DB *gorm.DB
}

func (s *ServerRun) Route() *mux.Router {
	ctx := context.Background()
	vehicleService := service.VehicleService{
		DB: s.DB,
		VehicleStore: &store.VehicleStore{
			Ctx: ctx,
			DB:  s.DB,
		},
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/vehicle/vins/{vid}", vehicleService.FindByID).Methods("GET")
	router.HandleFunc("/api/vehicle/vins", vehicleService.FindAll).Methods("GET").
		Queries("limit", "{limit}", "offset", "{offset}")
	router.HandleFunc("/api/vehicle/vin", vehicleService.Create).Methods("POST")
	router.HandleFunc("/api/vehicle/vins/{vid}", vehicleService.Delete).Methods("DELETE")
	return router
}
