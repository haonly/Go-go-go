package service

import (
	"Go-go-go/ParkJiHyo/goAPI/domain/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type vehicleStoreFactory func(db *gorm.DB) model.VehicleStore

type VehicleService struct {
	DB           *gorm.DB
	VehicleStore model.VehicleStore
}

func (s *VehicleService) FindByID(w http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	response, err := s.VehicleStore.FindByID(vars["vid"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJson(w, http.StatusOK, response)
}
func (s *VehicleService) FindAll(w http.ResponseWriter, request *http.Request) {
	limit, _ := strconv.Atoi(request.FormValue("limit"))
	offset, _ := strconv.Atoi(request.FormValue("offset"))

	response, err := s.VehicleStore.FindAll(limit, offset)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJson(w, http.StatusOK, response)
}
func (s *VehicleService) Create(w http.ResponseWriter, request *http.Request) {
	tx := s.DB.Begin()
	if tx.Error != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}
	defer tx.RollbackUnlessCommitted()

	vehicle := model.VehicleEntity{}
	if err := json.NewDecoder(request.Body).Decode(&vehicle); err != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}

	if err := s.VehicleStore.Create(vehicle); err != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}

	if err := tx.Commit().Error; err != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}
	respondWithJson(w, http.StatusOK, vehicle)
}
func (s *VehicleService) Delete(w http.ResponseWriter, request *http.Request) {
	tx := s.DB.Begin()
	if tx.Error != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}
	defer tx.RollbackUnlessCommitted()

	vars := mux.Vars(request)
	if err := s.VehicleStore.Delete(vars["vid"]); err != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}

	if err := tx.Commit().Error; err != nil {
		respondWithError(w, http.StatusBadRequest, tx.Error.Error())
	}
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
