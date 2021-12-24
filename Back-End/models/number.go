package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Number struct {
	gorm.Model
	Numverify datatypes.JSON
	Google  datatypes.JSON
	SearchID uint
	Search Search
}

func CreatePhoneRecord(number Number, db *gorm.DB) error  {


	if err := db.Create(&number).Error; err != nil {
		return err
	}
	return nil
}
