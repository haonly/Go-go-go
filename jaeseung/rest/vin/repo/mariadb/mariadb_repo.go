package mariadb

import (
	"context"
	"fmt"
	mysql "go.elastic.co/apm/module/apmgormv2/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config interface {
	DBAddress() string
	DBPort() int
	DBName() string
	DBUser() string
	DBPassword() string
}

type MariaDBRepository struct {
	db *gorm.DB
}

type VIN struct {
	Vin string `gorm:"primarykey"`
	Vid string
}

func New(c Config) *MariaDBRepository {
	//dsn := "root:1234@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := ""
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser(),
		c.DBPassword(),
		c.DBAddress(),
		c.DBPort(),
		c.DBName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&VIN{}); err != nil {
		log.Fatal(err)
	}
	return &MariaDBRepository{db: db}
}

func (m *MariaDBRepository) Save(ctx context.Context, vin string, vid string) (string, error) {
	m.db = m.db.WithContext(ctx)
	v := VIN{vin, vid}
	result := m.db.Create(&v) // pass pointer of data to Create

	if result.Error != nil {
		return "", result.Error
	}
	return v.Vid, nil
}

func (m *MariaDBRepository) Get(ctx context.Context, vin string) (vid string, e error) {
	m.db = m.db.WithContext(ctx)
	v := &VIN{}
	err := m.db.Where("vin=?", vin).First(v).Error
	if err != nil {
		return "", fmt.Errorf("DB Get error=%s", err)
	}
	return v.Vid, nil
}

func (m *MariaDBRepository) Update(ctx context.Context, vin string, vid string) error {
	m.db = m.db.WithContext(ctx)
	v := &VIN{Vid: vid}
	err := m.db.Update("vid", v).Where("vin=?", vin).Error
	if err != nil {
		return fmt.Errorf("DB Update error=%s", err)
	}
	return nil
}

func (m *MariaDBRepository) Delete(ctx context.Context, vin string) (string, error) {
	m.db = m.db.WithContext(ctx)
	v := VIN{Vin: vin}
	vid, err := m.Get(ctx, vin)
	if err != nil {
		return "", err
	}
	m.db.Delete(&v)
	return vid, nil
}
