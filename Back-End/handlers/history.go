package handlers

import (
	"back-end/database"
	"back-end/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

func GetMyHistory(c *fiber.Ctx) error  {

	type SearchHistory struct {

		Keyword string `json:"keyword"`
		Offset int `json:"offset"`
		Limit int `json:"limit"`
		Action string `json:"action"`

	}

	historySearch := new(SearchHistory)
	c.Accepts("application/json")


	err := c.BodyParser(historySearch )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}




	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	// var search []models.Search
	 var result []map[string]interface{}
	// database.DBInstance.Select("created_at","s_type","keyword").Find(&models.Search{} , "user_id = ?" , uid)
	var query *gorm.DB
	var total int64 = 0

	if historySearch.Action != "" {

		// get total history count
		var count int64
		database.DBInstance.Model(&models.Search{}).
			Select("created_at","s_type","keyword").
			Where("keyword LIKE ?", "%"+historySearch.Keyword+"%").
			Or("s_type LIKE ?", "%"+historySearch.Keyword+"%").
			Where("user_id = ?" , uid).Count(&count)
		total = count


		query = database.DBInstance.Model(&models.Search{}).
			Select("id","created_at", "s_type", "keyword").
			Limit(historySearch.Limit).Offset(historySearch.Offset).
			Order("created_at DESC").
			Where("keyword LIKE ?", "%"+historySearch.Keyword+"%").
			Or("s_type LIKE ?", "%"+historySearch.Keyword+"%").
			Find(&result, "user_id = ?", uid)

	}else {


		// get total history count
		var count int64
		database.DBInstance.Model(&models.Search{}).
			Select("created_at","s_type","keyword").
			Where("user_id = ?" , uid).Count(&count)
		total = count


		query = database.DBInstance.Model(&models.Search{}).
			Select("id","created_at","s_type","keyword").
			Limit(historySearch.Limit).Offset(historySearch.Offset).
			Order("created_at DESC").
			Find(&result,"user_id = ?" , uid)

	}

	if query.Error != nil{

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// testing loader
	// time.Sleep(300 * time.Millisecond)


	return c.JSON(fiber.Map{"history":result , "total":total })

}

func GetHistoryByID(c *fiber.Ctx)  error {


	type SearchID struct {
		ID string `json:"search_id"`
		SType string `json:"s_type"`
	}


	search := new(SearchID)
	c.Accepts("application/json") // "application/json"

	//log.Println(string(c.Body()))




	err := c.BodyParser(search )
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}




	// get user ID
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))


	var historyIDS []uint
	result := database.DBInstance.Model(&models.Search{}).Pluck("user_id",&historyIDS)

	if result.Error != nil{

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if result.RowsAffected < 1 {
		return c.SendStatus(fiber.StatusNotFound)
	}


	// IDOR check ;)
	access := false
	for _,id := range historyIDS {

		if id == uint(uid) {
			access  = true
			break
		}
	}
	if !access {
		return c.SendStatus(fiber.StatusUnauthorized)

	}



	if  search.SType == "number"{

		var number models.Number
		phoneResult := database.DBInstance.First(&number ,  "search_id = ?" , search.ID)
		if phoneResult.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(fiber.Map{"numverify":number.Numverify , "google":number.Google , "searchID" :search.ID})


	}

	if  search.SType == "ip"{

		var ipModel models.Ip
		ipResult := database.DBInstance.First(&ipModel ,  "search_id = ?" , search.ID)
		if ipResult.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(fiber.Map{ "searchID" : search.ID ,  "ipData" : ipModel.IpData , "whois" :ipModel.Whois})


	}


	if  search.SType == "domain"{

		var domain models.Domain
		domainResult := database.DBInstance.First(&domain ,  "search_id = ?" , search.ID)
		if domainResult.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		screenShots :=  strings.Split(domain.Screenshots, "||")

		return c.JSON(fiber.Map{ "searchID" :search.ID  , "ipData" : domain.IpData , "whois" :domain.Whois , "screenShots" : screenShots} )


	}

	if  search.SType == "username"{

		var username models.Username
		userResult := database.DBInstance.First(&username ,  "search_id = ?" , search.ID)
		if userResult.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(fiber.Map{ "searchID" :search.ID , "socialAccounts" : username.Accounts})

	}




	if  search.SType == "association"{




		var leak models.Dehashed
		associationResults := database.DBInstance.First(&leak ,  "search_id = ?" , search.ID)
		if associationResults.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(fiber.Map{"emails":leak.Emails , "usernames":leak.Usernames,
			"ips":leak.Ips , "numbers":leak.Numbers , "names":leak.Names , "addresses":leak.Addresses, "domains":leak.Domains })


	}


	return nil

}

