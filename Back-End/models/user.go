package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name       string `gorm:"not null" `
	Email      string `gorm:"unique;not null;"`
	VerifiedAT time.Time `gorm:"default:null"`
	Password   string `gorm:"not null;"`
	Credit     int32 `gorm:"default:100"`
	UType      int32 `gorm:"default:2"`
	search []Search
}



func GetUserByEmail( email string , db *gorm.DB) (User, error) {
	var user User
	//db := database.DBInstance
	res := db.First(&user , "email = ?" , email)
	if res.Error != nil {
		return User{} , res.Error
	}
	return user , nil

}
func CreateUser(user User   , db *gorm.DB) error  {

	//db := database.DBInstance
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil

}


func EmailExists(email string , db *gorm.DB)  bool {

	//db := database.DBInstance
	var user User

	//.Debug()
	res := db.Find(&user, "email = ?" , email)

	// anything ?
	if res.RowsAffected > 0 {
		return true
	}
	return false

}

