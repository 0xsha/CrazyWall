package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)


type Dehashed struct {
	gorm.Model
	Emails datatypes.JSON
	Names datatypes.JSON
	Domains datatypes.JSON
	Ips datatypes.JSON
	Usernames datatypes.JSON
	Addresses datatypes.JSON
	Numbers datatypes.JSON
	SearchID uint
	Search Search
}

func CreateLeakData(dehashed Dehashed, db *gorm.DB) error  {


	if err := db.Create(&dehashed).Error; err != nil {
		return err
	}
	return nil

}