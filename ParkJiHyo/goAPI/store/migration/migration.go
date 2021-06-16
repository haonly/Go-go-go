package migration

import (
	"Go-go-go/ParkJiHyo/goAPI/store/datamodel"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Do(db *gorm.DB) error {
	if err := migrateVehicle(db); err != nil {
		return err
	}
	return nil
}

func migrateVehicle(db *gorm.DB) error {
	err := db.AutoMigrate(&datamodel.Vehicle{}).Error
	if err != nil {
		fmt.Printf("mirgate vehicle error")
	}
	return err
}
