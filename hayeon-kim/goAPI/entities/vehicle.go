package entities

import "fmt"

type Vehicle struct {
	ID 			int64		`json:"ID"`
	VIN 		string	`json:"VIN"`
	VID 		string	`json:"VID"`
	CREATE_TIME string	`json:"CREATE_TIME"`
	MODIFY_TIME string	`json:"MODIFY_TIME"`
}


type ModifyVehicle struct {
	ID 			int64		`json:"ID"`
	VIN 		string	`json:"VIN"`
}



func (vehicle Vehicle) ToString() string{
	return fmt.Sprintf("VIN: %s\nVID: %s\bid, " +
		"create_time, modify_time are: %d, %s, %s",
		vehicle.VIN, vehicle.VID, vehicle.ID, vehicle.CREATE_TIME, vehicle.MODIFY_TIME)
}
