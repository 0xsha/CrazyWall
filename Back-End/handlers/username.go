package handlers

import (
	"back-end/database"
	"back-end/models"
	"back-end/modules"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func SearchUserName(c *fiber.Ctx) error  {


	// memory will be freed after function ;)
	type Username struct {
		UserName string `json:"username"`

	}


	userName := new(Username)
	c.Accepts("application/json") // "application/json"



	err := c.BodyParser(userName )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}



	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))



	// create the search record
	var search models.Search
	search.SType = "username"
	search.UserID = uint(uid)
	search.Keyword = userName.UserName


	socialAccounts, err := modules.GetSocialAccounts(userName.UserName)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	//log.Println(socialAccounts)



	id, err := models.CreateSearch(search, database.DBInstance)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}



	// create record data
	var username models.Username
	username.SearchID = id
	username.Accounts = socialAccounts.(string)
	//number.Google , _ = json.Marshal(resGoogle)



	err = models.CreateAccountsRecord(username,database.DBInstance)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}



	//accList := strings.Split(socialJson["output"], "\\n")
	//
	//for _, i := range(accList) {
	//	if strings.Index(i, ": ") != -1 {
	//
	//		indexColon := strings.Index(i, ": ")
	//		fmt.Println(i[indexColon+2:])
	//
	//	}
	//}

	return c.JSON(fiber.Map{ "searchID" :id , "socialAccounts" : socialAccounts})
}

