package vehicle

import (
	"Go-go-go/hayeon-kim/goAPI/config"
	"Go-go-go/hayeon-kim/goAPI/entities"
	"Go-go-go/hayeon-kim/goAPI/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		vehicleModel := models.VehicleModel{
			DB: db,
		}
		vehicles, err2 := vehicleModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, vehicles)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		vehicleModel := models.VehicleModel{
			DB: db,
		}
		vehicles, err2 := vehicleModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, vehicles)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var vehicle entities.ModifyVehicle
	err := json.NewDecoder(request.Body).Decode(&vehicle)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		vehicleModel := models.VehicleModel{
			DB: db,
		}
		err2 := vehicleModel.Create(&vehicle)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, vehicle)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	var vehicle entities.ModifyVehicle
	err := json.NewDecoder(request.Body).Decode(&vehicle)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		vehicleModel := models.VehicleModel{
			DB: db,
		}
		VIN, _, err2 := vehicleModel.Delete(&vehicle)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, map[string]string{
				"VinAffected": VIN,
			})
		}
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
