package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Ip struct {
	gorm.Model
	IpData datatypes.JSON
	Whois datatypes.JSON
	SearchID uint
	Search Search
}

func CreateIpRecord(ip Ip, db *gorm.DB) error  {


	if err := db.Create(&ip).Error; err != nil {
		return err
	}
	return nil
}

