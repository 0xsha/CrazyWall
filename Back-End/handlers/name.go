package handlers

import (
	"back-end/database"
	"back-end/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SearchName(c *fiber.Ctx) error  {


	// memory will be freed after function ;)
	type Name struct {
		Name string `json:"name"`

	}


	name := new(Name)
	c.Accepts("application/json") // "application/json"



	err := c.BodyParser(name )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}



	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))


	// create the search record
	var search models.Search
	search.SType = "name"
	search.UserID = uint(uid)
	search.Keyword = name.Name

	id, err := models.CreateSearch(search, database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{ "searchID" :id})
}

