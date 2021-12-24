package handlers

import (
	"back-end/config"
	"back-end/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetFlagURL(c *fiber.Ctx) error {

	type RestEU struct {
		Flag string `json:"flag"`
	}

	type CountryCode struct {
		Code string `json:"code"`
	}

	f := new(RestEU)

	code := new(CountryCode)
	c.Accepts("application/json") // "application/json"

	err := c.BodyParser(code)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	url := fmt.Sprintf("https://"+config.GetKey("RESETCOUNTIES_URL")+"/rest/v2/alpha/%s", code.Code)

	flag, err := helpers.GetJson(url, f, "", "")
	if err != nil {
		return err
	}

	return c.JSON(flag)

}
