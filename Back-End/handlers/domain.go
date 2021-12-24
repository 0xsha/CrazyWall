package handlers

import (
	"back-end/database"
	"back-end/helpers"
	"back-end/models"
	"back-end/modules"
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func SearchDomain(c *fiber.Ctx) error  {


	// memory will be freed after function ;)
	type Domain struct {
		Domain string `json:"domain"`

	}


	domain := new(Domain)
	c.Accepts("application/json") // "application/json"




	err := c.BodyParser(domain )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}




	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))


	// create the search record
	var search models.Search
	search.SType = "domain"
	search.UserID = uint(uid)
	search.Keyword = domain.Domain


	id, err := models.CreateSearch(search, database.DBInstance)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}


	whois, err := modules.GetWhois(domain.Domain)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	IP , err := helpers.GetIPByDomain(domain.Domain)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//
	ipData, err := modules.GetIpData(IP)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	screenShots, err := modules.GetScreenshots(domain.Domain)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}



	var domainModel models.Domain


	domainModel.SearchID = id
	domainModel.Whois, _ = json.Marshal(whois)
	domainModel.IpData, _ = json.Marshal(ipData)

	flattenShots := strings.Join(screenShots,"||")

	domainModel.Screenshots = flattenShots

	err = models.CreateDomainRecord(domainModel,database.DBInstance)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}


	return c.JSON(fiber.Map{ "searchID" :id  , "ipData" : ipData , "whois" :whois , "screenShots" : screenShots} )
}

