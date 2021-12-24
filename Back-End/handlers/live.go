package handlers

import (
	//"back-end/database"
	//"back-end/models"
	//"back-end/modules"
	//"encoding/json"
	//"github.com/form3tech-oss/jwt-go"
	"back-end/modules"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SearchLive(c *fiber.Ctx) error {

	// memory will be freed after function ;)
	type Keyword struct {
		Keyword string `json:"keyword"`
	}




	c.Accepts("application/json") // "application/json"


	keyword := new(Keyword)


	err := c.BodyParser(keyword)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}


	newsData, err := modules.GetNews(keyword.Keyword)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Println("========================================================================================================")
	socialData, err := modules.GetSocial(keyword.Keyword)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}




	// get user ID
	//token := c.Locals("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//uid := int(claims["user_id"].(float64))
	//
	//// create the search record
	//var search models.Search
	//search.SType = "ip"
	//search.UserID = uint(uid)
	//search.Keyword = ip.IP
	//
	//id, err := models.CreateSearch(search, database.DBInstance)
	//if err != nil {
	//	return c.SendStatus(fiber.StatusInternalServerError)
	//}

	return c.JSON(fiber.Map{ "news" : newsData  , "social" : socialData})



}
