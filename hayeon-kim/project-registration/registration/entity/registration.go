package entity

import (
	"github.com/google/uuid"
	"time"
)

type Registration struct {
	Id 				string `json:"Id"`
	VehicleID		string `json:"vehicleID"`
	MasterUserID	string `json:"masterUserID"`
	Company 		string `json:"Company"`
	Type			string `json:"type"`
	Year			int    `json:"year"`
	Capacity		int	   `json:"capacity"`
	CreateTime		int64  `json:"create_time"`
	ModifyTime		int64  `json:"modify_time"`
}

func NewRegistration(masterUserID string) * Registration {
	return &Registration{
		Id:				uuid.New().String(),
		VehicleID:		uuid.New().String(),
		MasterUserID :	masterUserID,
		CreateTime:		time.Now().UTC().Unix(),
	}
}

func (r *Registration) WithCompany(company string) *Registration {
	r.Company = company
	return r
}

func (r *Registration) WithType(_type string) *Registration{
	r.Type = _type
	return r
}

func (r *Registration) WithYear(year int) *Registration{
	r.Year = year
	return r
}

func (r *Registration) WithCapacity(capacity int) *Registration{
	r.Capacity = capacity
	return r
}