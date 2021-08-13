package registration

import (
	"net/http"
)

func Register(response http.ResponseWriter, request *http.Request) {
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
