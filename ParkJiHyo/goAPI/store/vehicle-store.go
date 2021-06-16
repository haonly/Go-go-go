package store

import (
	"Go-go-go/ParkJiHyo/goAPI/domain/model"
	"Go-go-go/ParkJiHyo/goAPI/store/datamodel"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type VehicleStore struct {
	Ctx context.Context
	DB  *gorm.DB
}

func (s *VehicleStore) FindByID(vid string) (*model.VehicleEntity, error) {
	vehicle := datamodel.Vehicle{}
	db := s.DB.Model(vehicle)

	if err := db.Where(&datamodel.Vehicle{VID: vid}).First(&vehicle).Error; err != nil {
		return nil, errorMsg(fmt.Sprintln("[DB Error] ", err))
	}
	return &model.VehicleEntity{VID: vehicle.VID, VIN: vehicle.VIN}, nil
}

func (s *VehicleStore) FindAll(limit, offset int) ([]*model.VehicleEntity, error) {
	var vehicles []datamodel.Vehicle
	db := s.DB.Model(&datamodel.Vehicle{})
	db = db.Offset(offset).Limit(limit)
	if err := db.Scan(&vehicles).Error; err != nil {
		return nil, errorMsg(fmt.Sprintln("[DB Error] ", err))
	}

	vehicleEntities := make([]*model.VehicleEntity, len(vehicles))
	for i := 0; i < len(vehicles); i++ {
		vehicleEntities[i] = &model.VehicleEntity{
			VIN: vehicles[i].VIN,
			VID: vehicles[i].VID,
		}
	}
	return vehicleEntities, nil
}

func (s *VehicleStore) Create(entity model.VehicleEntity) error {
	vid := uuid.New()
	vehicle := datamodel.Vehicle{VID: vid.String(), VIN: entity.VIN}
	if err := s.DB.Create(&vehicle).Error; err != nil {
		return errorMsg(fmt.Sprintln("[DB Error] ", err))
	}
	return nil
}

func (s *VehicleStore) Delete(vid string) error {
	if err := s.DB.Where(&datamodel.Vehicle{VID: vid}).Delete(&datamodel.Vehicle{}).Error; err != nil {
		return errorMsg(fmt.Sprintln("[DB Error] ", err))
	}
	return nil
}

func errorMsg(msg string) error {
	return errors.New(msg)
}
