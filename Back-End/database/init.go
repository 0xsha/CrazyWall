package database

import (
	"back-end/config"
	"back-end/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB()  {


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local" ,
		config.GetKey("DATABASE_USER") , config.GetKey("DATABASE_PASSWORD") ,
		config.GetKey("DATABASE_HOST")  , config.GetKey("DATABASE_PORT") , config.GetKey("DATABASE_NAME")  )

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}


	err = db.AutoMigrate(&models.User{},&models.Search{},&models.Dehashed{},&models.Number{} , &models.Ip{}, &models.Username{} , &models.Domain{})
	if err != nil {
		panic("failed to migrate")
	}

	// set global instance
	DBInstance = db



}

