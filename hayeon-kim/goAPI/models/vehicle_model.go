package models

import (
	"Go-go-go/hayeon-kim/goAPI/entities"
	"database/sql"
	"fmt"
	"hash/fnv"
)

type VehicleModel struct {
	DB *sql.DB
}

func (vehicleModel VehicleModel) FindAll() (vehicle []entities.Vehicle, error error) {
	rows, err := vehicleModel.DB.Query("Select * from vehicle;")
	return returnVehicleValuesInFormat(rows, err)
}

func (vehicleModel VehicleModel) Search(keyword string) (vehicle []entities.Vehicle, error error) {
	rows, err := vehicleModel.DB.Query("Select * from vehicle where VIN like ?", "%"+keyword+"%")
	return returnVehicleValuesInFormat(rows, err)
}

func (vehicleModel VehicleModel) Create(getVehicle *entities.ModifyVehicle) (error error) {
	VID := hash(getVehicle.VIN)
	fmt.Println("VIN, VID: "+getVehicle.VIN, VID)
	result, err := vehicleModel.DB.Exec("insert into vehicle (VIN, VID, create_time, modify_time) values (?, ?, now(), now());", getVehicle.VIN, VID)
	if err != nil {
		return err
	} else {
		getVehicle.ID, _ = result.LastInsertId()
		return nil
	}
}

func (vehicleModel VehicleModel) Delete(getVehicle *entities.ModifyVehicle) (string, int64, error) {
	VIN := getVehicle.VIN
	result, err := vehicleModel.DB.Exec("delete from vehicle where VIN=?", VIN)
	if err != nil {
		return VIN, 0, err
	} else {
		ID, retError := result.RowsAffected()
		getVehicle.ID = ID
		return VIN, getVehicle.ID, retError
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func returnVehicleValuesInFormat(rows *sql.Rows, err error) (vehicle []entities.Vehicle, error error) {
	if err != nil {
		return nil, err
	} else {
		var vehicles []entities.Vehicle
		for rows.Next() {
			var id int64
			var vin string
			var vid string
			var create_time string
			var modify_time string
			err2 := rows.Scan(&id, &vin, &vid, &create_time, &modify_time)
			if err2 != nil {
				return nil, err2
			} else {
				vehicle := entities.Vehicle{
					ID:          id,
					VIN:         vin,
					VID:         vid,
					CREATE_TIME: create_time,
					MODIFY_TIME: modify_time,
				}
				vehicles = append(vehicles, vehicle)
			}
		}
		return vehicles, nil
	}
}
