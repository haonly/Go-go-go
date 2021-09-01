package database

import (
	"Go-go-go/hayeon-kim/registration/app/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type RegistrationDB interface {
	Open() error
	Close() error
	CreateRegistration(p *models.Registration) error
	GetRegistration() ([]*models.Registration, error)
}

// TODO:
//db 인스턴스 설명
type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")

	pg.MustExec(createShema)

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
