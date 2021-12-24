package models

import (
	"gorm.io/gorm"
)

type Search struct {
	gorm.Model
	SType string
	UserID uint
	Keyword string
}

func CreateSearch(search Search , db *gorm.DB ) ( uint, error) {

	//db := database.DBInstance
	if err := db.Create(&search).Error; err != nil {
		return 0,err
	}
	return search.ID , nil

}