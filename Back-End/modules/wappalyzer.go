package modules

//import (
//	"back-end/config"
//	"back-end/helpers"
//	"fmt"
//)
//
//type wappalyzer []struct {
//	Technologies []struct {
//		Categories []struct {
//			ID   int64  `json:"id"`
//			Name string `json:"name"`
//			Slug string `json:"slug"`
//		} `json:"categories"`
//		Name     string   `json:"name"`
//		Slug     string   `json:"slug"`
//		Versions []string `json:"versions"`
//	} `json:"technologies"`
//	URL string `json:"url"`
//}
//
//func GetTechStack(domain string) (interface{} , error)  {
//
//	endpoint := config.GetKey("WAPPALYZER_ENDPOINT")
//	wapPath :=  fmt.Sprintf("/live/v2/?urls=%s&recursive=false" , domain)
//
//	fullUrl := "https://"+endpoint+wapPath
//	wObj := new(reddit)
//
//	_, err := helpers.GetJson(fullUrl, wObj , "wappalyzer",config.GetKey("WAPPALYZER_AUTH"))
//	if err != nil {
//		return nil, err
//	}
//
//	return wObj , nil
//
//}