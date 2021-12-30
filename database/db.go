package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instanceSingleTon *gorm.DB

func GetInstance() *gorm.DB {
	return instanceSingleTon
}

func init() {
	var err error
	instanceSingleTon, err = gorm.Open(sqlite.Open("./smart-gateway.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
}
