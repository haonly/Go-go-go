package model

type VehicleStore interface {
	FindByID(vid string) (*VehicleEntity, error)
	FindAll(limit, offset int) ([]*VehicleEntity, error)
	Create(model VehicleEntity) error
	Delete(vid string) error
}
