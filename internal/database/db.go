package database

import (
	"grpc/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB{
	dsn:="host=localhost user=postgres password=A1barsoap dbname=dbgrpc port=5432 sslmode=disable"

	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	return db
}