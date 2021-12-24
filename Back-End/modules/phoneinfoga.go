package modules

import (
	"back-end/config"
	"back-end/helpers"
	"fmt"
)

type numverifyObj struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Result  struct {
		Valid               bool   `json:"valid"`
		Number              string `json:"number"`
		LocalFormat         string `json:"local_format"`
		InternationalFormat string `json:"international_format"`
		CountryPrefix       string `json:"country_prefix"`
		CountryCode         string `json:"country_code"`
		CountryName         string `json:"country_name"`
		Location            string `json:"location"`
		Carrier             string `json:"carrier"`
		LineType            string `json:"line_type"`
	} `json:"result"`
}


type GoogleSearch struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Result  struct {
		Socialmedia []struct {
			Number string `json:"number"`
			Dork   string `json:"dork"`
			URL    string `json:"URL"`
		} `json:"socialMedia"`
		Disposableproviders []struct {
			Number string `json:"number"`
			Dork   string `json:"dork"`
			URL    string `json:"URL"`
		} `json:"disposableProviders"`
		Reputation []struct {
			Number string `json:"number"`
			Dork   string `json:"dork"`
			URL    string `json:"URL"`
		} `json:"reputation"`
		Individuals []struct {
			Number string `json:"number"`
			Dork   string `json:"dork"`
			URL    string `json:"URL"`
		} `json:"individuals"`
		General []struct {
			Number string `json:"number"`
			Dork   string `json:"dork"`
			URL    string `json:"URL"`
		} `json:"general"`
	} `json:"result"`
}

func GetNumberInfo(number string) (interface{}, interface{}, error) {



	endpoint := config.GetKey("PHONEINFOGA_URL")

	numverify :=  fmt.Sprintf("/api/numbers/%s/scan/numverify" , number)
	google :=  fmt.Sprintf("/api/numbers/%s/scan/googlesearch" , number)

	//ovh :=  fmt.Sprintf("api/numbers/%s/scan/ovh" , number)
	//isValid :=  fmt.Sprintf("api/numbers/%s/validate" , number)

	fullUrl := "https://"+endpoint+numverify
	nObj := new(numverifyObj)


	_, err := helpers.GetJson(fullUrl, nObj , "","")
	if err != nil {
		return nil,nil, err
	}
	fullUrl = "https://"+endpoint+google
	gObj := new(GoogleSearch)

	_, err = helpers.GetJson(fullUrl, gObj , "" ,"")
	if err != nil {
		return nil,nil, err
	}

	//log.Println(gObj)
	//log.Println(nObj)

	return nObj,gObj,nil
}
