package handlers


import (
	"back-end/database"
	"back-end/models"
	"back-end/modules"
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SearchIP(c *fiber.Ctx) error  {


	// memory will be freed after function ;)
	type IP struct {
		IP string `json:"ip"`

	}


	ip := new(IP)
	c.Accepts("application/json") // "application/json"



	err := c.BodyParser(ip )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}



	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))


	// create the search record
	var search models.Search
	search.SType = "ip"
	search.UserID = uint(uid)
	search.Keyword = ip.IP

	id, err := models.CreateSearch(search, database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	//
	ipData, err := modules.GetIpData(ip.IP)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	whois, err := modules.GetWhois(ip.IP)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	var iPmodel models.Ip

	iPmodel.SearchID = id
	iPmodel.Whois, _ = json.Marshal(whois)
	iPmodel.IpData, _ = json.Marshal(ipData)

	err = models.CreateIpRecord(iPmodel,database.DBInstance)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}


	return c.JSON(fiber.Map{ "searchID" :id ,  "ipData" : ipData , "whois" :whois} )
}
