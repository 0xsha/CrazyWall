package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Screenshots string
	//Subdomains datatypes.JSON
	IpData   datatypes.JSON
	Whois    datatypes.JSON
	SearchID uint
	Search   Search
}



func CreateDomainRecord(domain Domain, db *gorm.DB) error  {


	if err := db.Create(&domain).Error; err != nil {
		return err
	}
	return nil
}


