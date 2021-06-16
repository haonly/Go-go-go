package datamodel

import "github.com/jinzhu/gorm"

type Vehicle struct {
	gorm.Model
	VID string
	VIN string
}

func (Vehicle) TableName() string {
	return "vehicle"
}
