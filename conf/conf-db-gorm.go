package conf

import (
	"github.com/Williberto14/powercrud/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func ConfDatabase() *gorm.DB {

	dsn := "root:qwerty@tcp(127.0.0.1:3310)/prueba_tecnica?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	errMigrate := db.AutoMigrate(&model.Product{})
	if errMigrate != nil {
		panic(errMigrate)
	}

	return db
}
