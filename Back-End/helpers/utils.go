package helpers

import (
	"net"
)

func Unique(input []string) []string {
	unique := make(map[string]bool, len(input))
	list := make([]string, len(unique))
	for _, el := range input {
		if len(el) != 0 {
			if !unique[el] {
				list = append(list, el)
				unique[el] = true
			}
		}
	}
	return list
}

func GetIPByDomain(domain string) (string  , error) {

	lookupHost, err := net.LookupHost(domain)
	if err != nil {
		return "", err
	}
	return lookupHost[0] , nil

	// online version
	//type IpData struct {
	//	Query       string  `json:"query"`
	//}
	//
	//ipData := new(IpData)
	//
	//json, err := GetJson("http://ip-api.com/json/"+domain, ipData, "", "")
	//if err != nil {
	//	return "" , err
	//}
	//
	//return json.(string) , nil

}


