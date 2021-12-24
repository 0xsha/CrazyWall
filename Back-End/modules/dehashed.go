package modules

import (
	"back-end/config"
	"back-end/helpers"
	"fmt"
	"strings"
)

type Dehashed struct {
	Balance int `json:"balance"`
	Entries []struct {
		ID             string `json:"id"`
		Email          string `json:"email"`
		Username       string `json:"username"`
		Password       string `json:"password"`
		HashedPassword string `json:"hashed_password"`
		Name           string `json:"name"`
		Vin            string `json:"vin"`
		Address        string `json:"address"`
		IPAddress      string `json:"ip_address"`
		Phone          string `json:"phone"`
		ObtainedFrom   string `json:"database_name"`
	} `json:"entries"`
	Success bool   `json:"success"`
	Took    string `json:"took"`
	Total   int    `json:"total"`
}

// GetDehashedData we limits only 1k results for now
func GetDehashedData(iType string, input string) (interface{}, error) {

	authCred := config.GetKey("DEHASHED_AUTH")
	baseUrl := config.GetKey("DEHASHED_ENDPOINT")
	pathUrl := ""



	switch iType {

	// only search 'exact'
	//domain:example.com
	//username:example.com
	//email:example.com
	//name:
	//username:test&size=50

	case "email":
		pathUrl =  fmt.Sprintf("/search?query=email:'%s'&size=1000" , input)
	case "domain":
		pathUrl =  fmt.Sprintf("/search?query=domain:'%s'&size=1000" , input)
	case "username":
		pathUrl =  fmt.Sprintf("/search?query=username:'%s'&size=1000" , input)
	case "number":
		pathUrl =  fmt.Sprintf("/search?query=phone:'%s'&size=1000" , input)
	case "name":
		pathUrl =  fmt.Sprintf("/search?query=name:'%s'" , strings.Replace(input," ","%2b", -1))
	case "ip":
		pathUrl =  fmt.Sprintf("/search?query=ip_address:'%s'" , input)


	}

	fullUrl := "https://"+ baseUrl + pathUrl


	dehashedObj := new (Dehashed)

	dObj, err := helpers.GetJson(fullUrl,dehashedObj,"dehashed",authCred)
	if err != nil {

		return nil, err
	}


	return dObj,nil



}

