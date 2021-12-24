package models

import (
	"gorm.io/gorm"
)

type Username struct {
	gorm.Model
	//Accounts datatypes.JSON
	SearchID uint
	Search   Search
	Accounts string

}


func CreateAccountsRecord(username Username, db *gorm.DB) error  {


	if err := db.Create(&username).Error; err != nil {
		return err
	}
	return nil
}

