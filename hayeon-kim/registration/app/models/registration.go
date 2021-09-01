package models

import (
	"time"
)

type Registration struct {
	Id           string    `db:"id"`
	VehicleID    string    `db:"vehicleID"`
	MasterUserID string    `db:"masterUserID"`
	Company      string    `db:"company"`
	Model        string    `db:"model"`
	Type         string    `db:"type"`
	Year         int       `db:"year"`
	Capacity     int       `db:"capacity"`
	CreateTime   time.Time `db:"create_time"`
	ModifyTime   time.Time `db:"modify_time"`
}

type JsonRegistration struct {
	Id           string    `json:"id"`
	VehicleID    string    `json:"vehicleID"`
	MasterUserID string    `json:"masterUserID"`
	Company      string    `json:"company"`
	Model        string    `json:"model"`
	Type         string    `json:"type"`
	Year         int       `json:"year"`
	Capacity     int       `json:"capacity"`
	CreateTime   time.Time `json:"create_time"`
	ModifyTime   time.Time `json:"modify_time"`
}

type RegistrationRequest struct {
	MasterUserID string `json:"masterUserID"`
	Company      string `json:"company"`
	Model        string `json:"model"`
	Type         string `json:"type"`
	Year         int    `json:"year"`
	Capacity     int    `json:"capacity"`
}
