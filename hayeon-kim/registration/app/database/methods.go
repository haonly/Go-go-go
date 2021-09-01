package database

import "Go-go-go/hayeon-kim/registration/app/models"

func (d *DB) CreateRegistration(p *models.Registration) error {
	res, err := d.db.Exec(insertRegistrationSchema, p.VehicleID, p.MasterUserID, p.Company, p.Model, p.Type, p.Year, p.Capacity)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetRegistration() ([]*models.Registration, error) {
	var registrations []*models.Registration
	err := d.db.Select(&registrations, "SELECT * FROM info_vehicle")
	if err != nil {
		return registrations, err
	}

	return registrations, nil
}
