package mariadb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MariaDBRepository struct {
	db *gorm.DB
}

type VIN struct {
	Vin string `gorm:"primarykey"`
	Vid string
}

func New() *MariaDBRepository {
	dsn := "root:1234@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&VIN{}); err != nil {
		log.Fatal(err)
	}
	return &MariaDBRepository{db: db}
}

func (m *MariaDBRepository) Save(vin string, vid string) (string, error) {
	v := VIN{vin, vid}
	result := m.db.Create(&v) // pass pointer of data to Create

	if result.Error != nil {
		return "", result.Error
	}
	return v.Vid, nil
}

func (m *MariaDBRepository) Get(vin string) (vid string, e error) {
	v := &VIN{}
	err := m.db.Where("vin=?", vin).First(v).Error
	if err != nil {
		return "", fmt.Errorf("DB Get error=%s", err)
	}
	return v.Vid, nil
}

func (m *MariaDBRepository) Update(vin string, vid string) error {
	v := &VIN{Vid: vid}
	err := m.db.Update("vid", v).Where("vin=?", vin).Error
	if err != nil {
		return fmt.Errorf("DB Update error=%s", err)
	}
	return nil
}

func (m *MariaDBRepository) Delete(vin string) (string, error) {
	v := VIN{Vin: vin}
	vid, err := m.Get(vin)
	if err != nil {
		return "", err
	}
	m.db.Delete(&v)
	return vid, nil
}
