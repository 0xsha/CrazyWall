package modules
//
//import (
//	"back-end/config"
//	"back-end/helpers"
//	"fmt"
//)
//
//
//type WAF struct {
//	Successful bool `json:"successful"`
//	Result string  `json:"result"`
//}
//
//func GetWaf(url string) (interface{} , error){
//
//
//	endpoint := config.GetKey("C99_ENDPOINT")
//
//	getPorts := fmt.Sprintf("/firewalldetector/?key=%s&url=%s&json" , config.GetKey("C99_AUTH") , url  )
//	fullUrl := "https://"+endpoint + getPorts
//
//	wObj := new(WAF)
//
//
//	_, err := helpers.GetJson(fullUrl, wObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return wObj , nil
//
//}
//
//type OpenPorts struct {
//	Port    string `json:"port"`
//	Success bool   `json:"success"`
//}
//
//func GetPorts(host string) (interface{} , error){
//
//
//	endpoint := config.GetKey("C99_ENDPOINT")
//
//	getPorts := fmt.Sprintf("/portscanner/?key=%s&host=%s&json" , config.GetKey("C99_AUTH") , host  )
//	fullUrl := "https://"+endpoint + getPorts
//
//	cObj := new(OpenPorts)
//
//
//	_, err := helpers.GetJson(fullUrl, cObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return cObj , nil
//
//
//}

////https://api.c99.nl/subdomainfinder?key=CYUOX-L783E-82BEA-3WNAQ&domain=c99.nl
//func GetSubDomains(){
// inCase
//}