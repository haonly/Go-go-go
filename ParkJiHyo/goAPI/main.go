package main

import (
	"Go-go-go/ParkJiHyo/goAPI/api"
	"Go-go-go/ParkJiHyo/goAPI/common/config"
	"Go-go-go/ParkJiHyo/goAPI/store/migration"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

func main() {
	// read environment
	//envFileName := "env.toml"
	//env := config.ReadEnv(envFileName)
	//log.Println(env)
	driverName, DataSourceName := config.GetDatabase()
	db := openDatabase(driverName, DataSourceName, gorm.Open)
	defer db.Close()
	router := api.ServerRun{DB: db}
	migrateDatabase(db)
	log.Print(http.ListenAndServe(fmt.Sprintf(":%d", 5000), router.Route()))
}

func openDatabase(driverName, DataSourceName string, openFunc func(string, ...interface{}) (*gorm.DB, error)) *gorm.DB {
	db, err := openFunc(driverName, DataSourceName)
	if err != nil {
		panic(err)
	}

	db.DB().SetConnMaxLifetime(time.Duration(30) * time.Second)
	db.DB().SetMaxOpenConns(10)
	return db
}
func migrateDatabase(db *gorm.DB) {
	err := migration.Do(db)
	if err != nil {
		panic(err)
	}
}
