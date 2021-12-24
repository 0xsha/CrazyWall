package handlers

import (
	"back-end/database"
	"back-end/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SearchEmail(c *fiber.Ctx) error  {


	// memory will be freed after function ;)
	type Email struct {
		Email string `json:"email"`

	}


	email := new(Email)
	c.Accepts("application/json") // "application/json"



	err := c.BodyParser(email )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}



	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))


	// create the search record
	var search models.Search
	search.SType = "email"
	search.UserID = uint(uid)
	search.Keyword = email.Email

	id, err := models.CreateSearch(search, database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{ "searchID" :id})
}

