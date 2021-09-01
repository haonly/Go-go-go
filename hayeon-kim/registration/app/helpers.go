package app

import (
	"Go-go-go/hayeon-kim/registration/app/models"
	"encoding/json"
	"log"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// json response를 같이 status에 보냄
func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

//json 변환
func mapRegistrationToJSON(r *models.Registration) models.JsonRegistration {
	return models.JsonRegistration{
		Id:           r.Id,
		VehicleID:    r.VehicleID,
		MasterUserID: r.MasterUserID,
		Company:      r.Company,
		Model:        r.Model,
		Type:         r.Type,
		Year:         r.Year,
		Capacity:     r.Capacity,
		CreateTime:   r.CreateTime,
		ModifyTime:   r.ModifyTime,
	}
}
