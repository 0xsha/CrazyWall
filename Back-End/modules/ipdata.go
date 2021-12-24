package modules

import (
	"back-end/config"
	"back-end/helpers"
	"fmt"
)

type IpDataCo struct {
	Asn struct {
		Asn    string `json:"asn"`
		Domain string `json:"domain"`
		Name   string `json:"name"`
		Route  string `json:"route"`
		Type   string `json:"type"`
	} `json:"asn"`
	CallingCode   string      `json:"calling_code"`
	City          interface{} `json:"city"`
	ContinentCode string      `json:"continent_code"`
	ContinentName string      `json:"continent_name"`
	Count         string      `json:"count"`
	CountryCode   string      `json:"country_code"`
	CountryName   string      `json:"country_name"`
	Currency      struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Native string `json:"native"`
		Plural string `json:"plural"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	EmojiFlag    string `json:"emoji_flag"`
	EmojiUnicode string `json:"emoji_unicode"`
	Flag         string `json:"flag"`
	IP           string `json:"ip"`
	IsEu         bool   `json:"is_eu"`
	Languages    []struct {
		Name   string `json:"name"`
		Native string `json:"native"`
	} `json:"languages"`
	Latitude   float64     `json:"latitude"`
	Longitude  float64     `json:"longitude"`
	Postal     interface{} `json:"postal"`
	Region     interface{} `json:"region"`
	RegionCode interface{} `json:"region_code"`
	Threat     struct {
		IsAnonymous     bool `json:"is_anonymous"`
		IsBogon         bool `json:"is_bogon"`
		IsKnownAbuser   bool `json:"is_known_abuser"`
		IsKnownAttacker bool `json:"is_known_attacker"`
		IsProxy         bool `json:"is_proxy"`
		IsThreat        bool `json:"is_threat"`
		IsTor           bool `json:"is_tor"`
	} `json:"threat"`
	TimeZone struct {
		Abbr        string `json:"abbr"`
		CurrentTime string `json:"current_time"`
		IsDst       bool   `json:"is_dst"`
		Name        string `json:"name"`
		Offset      string `json:"offset"`
	} `json:"time_zone"`
}

func GetIpData(ip string) (interface{} , error)  {


	endpoint := config.GetKey("IPDATA_ENDPOINT")

	ipdata :=  fmt.Sprintf("/%s?api-key=%s" , ip, config.GetKey("IPDATA_AUTH"))

	fullUrl := "https://"+endpoint+ipdata


	iObj := new(IpDataCo)

	_, err := helpers.GetJson(fullUrl, iObj , "","")
	if err != nil {
		return nil, err
	}

	return iObj , nil



}
