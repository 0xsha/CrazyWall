package handlers

import (
	"back-end/database"
	"back-end/models"
	"back-end/modules"
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SearchPhone(c *fiber.Ctx) error  {


	// memory will be freed after function ;)
	type Phone struct {
		Number string `json:"phone"`

	}


	phone := new(Phone)
	c.Accepts("application/json") // "application/json"



	err := c.BodyParser(phone )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}



	resNum,resGoogle,err := modules.GetNumberInfo(phone.Number)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)

	}


	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))


	// create the search record
	var search models.Search
	search.SType = "phone"
	search.UserID = uint(uid)
	search.Keyword = phone.Number

	id, err := models.CreateSearch(search, database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// create record data
	var number models.Number
	number.SearchID = id
	number.Numverify, _ = json.Marshal(resNum)
	number.Google , _ = json.Marshal(resGoogle)



	err = models.CreatePhoneRecord(number,database.DBInstance)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}


	return c.JSON(fiber.Map{"numverify":resNum , "google":resGoogle , "searchID" :id})
}

