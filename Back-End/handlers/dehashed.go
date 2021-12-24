package handlers

import (
	"back-end/database"
	"back-end/helpers"
	"back-end/models"
	"back-end/modules"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func SearchAssociation(c *fiber.Ctx) error {

	// memory will be freed after function ;)
	type SearchDB struct {
		Type string `json:"type"`
		Keyword string `json:"keyword"`
		ID string `json:"search_id"`

	}

	dbSearch := new(SearchDB)
	c.Accepts("application/json")


	err := c.BodyParser(dbSearch )
	if err != nil {


		return c.SendStatus(fiber.StatusBadRequest)
	}




	data, err := modules.GetDehashedData(dbSearch.Type, dbSearch.Keyword)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}



	// type assertion
	association := data.(*modules.Dehashed)

	//var passwords []string
	var (
		emails , ips, domains, addresses, names, numbers, usernames  []string
	)


	for _, entity := range association.Entries {

		//switch ent := entity; {
		//
		//case ent.Email != "":
		//	emails = append(emails, ent.Email)
		//
		//case ent.Name != "":
		//	names = append(names, ent.Name)
		//
		//case ent.ObtainedFrom != "":
		//	domains = append(domains, ent.ObtainedFrom)
		//
		//case ent.Phone != "":
		//	numbers = append(numbers, ent.Phone)
		//
		//case ent.IPAddress != "":
		//	ips = append(ips, ent.IPAddress)
		//
		//case ent.Address != "":
		//	addresses = append(addresses, ent.Address)
		//
		//case ent.Username != "":
		//	usernames = append(usernames, ent.Username)
		//
		//
		//}

		if entity.Email != "" && entity.Email != "NULL"{
			emails = append(emails, entity.Email)
		}
		if entity.Name != "" && entity.Name != "NULL"{
			names = append(names, entity.Name)
		}
		if entity.ObtainedFrom != "" && entity.ObtainedFrom != "NULL"{
			domains = append(domains, entity.ObtainedFrom)
		}

		if entity.Address != "" && entity.Address != "NULL"{
			addresses = append(addresses, entity.Address)
		}

		if entity.Username != "" && entity.Username != "NULL" {
			usernames = append(usernames, entity.Username)
		}

		//if entity.Password != "" && entity.Password != "NULL" {
		//	passwords = append(passwords, entity.Username)
		//}

		if entity.Phone != "" && entity.Phone != "NULL"{
			numbers = append(numbers, entity.Phone)
		}

		if entity.IPAddress != "" && entity.IPAddress != "NULL"{
			ips = append(ips, entity.IPAddress)
		}


	}

	// ahem , interested in unique only
	emails = helpers.Unique(emails)
	names = helpers.Unique(names)
	domains = helpers.Unique(domains)
	addresses = helpers.Unique(addresses)
	numbers = helpers.Unique(numbers)
	usernames = helpers.Unique(usernames)
	//passwords = helpers.Unique(passwords)
	ips = helpers.Unique(ips)



	// create association record
	var leak models.Dehashed
	intID , _ := strconv.Atoi(dbSearch.ID)
	leak.SearchID = uint( intID)
	leak.Emails , _ =  json.Marshal(emails)
	leak.Ips , _ =  json.Marshal(ips)
	leak.Names , _ =  json.Marshal(names)
	leak.Usernames , _ =  json.Marshal(usernames)
	leak.Addresses , _ =  json.Marshal(addresses)
	leak.Domains , _ =  json.Marshal(domains)
	leak.Numbers , _ = json.Marshal(numbers)
	//leak.passwords , _ =  json.Marshal(password)
	err = models.CreateLeakData(leak,database.DBInstance)
	if err != nil {

		return c.SendStatus(fiber.StatusInternalServerError)
	}




	// return all associations to front end
	return c.JSON(fiber.Map{"emails":emails , "usernames":usernames,
			"ips":ips , "numbers":numbers , "names":names , "addresses":addresses, "domains":domains })


}
