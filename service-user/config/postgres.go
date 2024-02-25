package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitPostgresDatabase() {
	dsn := "host=host.docker.internal user=postgres password=admin dbname=hrga_api port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetPostgresDatabase() *gorm.DB {
	return db
}
