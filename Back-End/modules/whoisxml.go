package modules

import (
	"back-end/config"
	"back-end/helpers"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func GetScreenshots(domain string) ([]string , error)  {


	endpoint := config.GetKey("WHOISXML_SCREENSHOT_ENDPOINT")

	iPhoneUA := "Mozilla/5.0 (iPhone; CPU iPhone OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"

	mobileScreenPath := fmt.Sprintf("/api/v1?apiKey=%s&url=%s&credits=SA&ua=%s&imageOutputFormat=base64" , config.GetKey("WHOISXML_AUTH") , domain  , iPhoneUA )
	mobileFullUrl := "https://"+endpoint + mobileScreenPath

	desktopScreenPath := fmt.Sprintf("/api/v1?apiKey=%s&url=%s&credits=SA&imageOutputFormat=base64" , config.GetKey("WHOISXML_AUTH") , domain  )
	desktopFullUrl := "https://"+endpoint + desktopScreenPath


	client := fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	var results []string


	req.SetRequestURI(mobileFullUrl)
	if err := client.DoTimeout(req, res, 30*time.Second); err != nil {
		//log.Println(err)
		return nil,err
	}

	results = append(results,  string(res.Body()))


	req.SetRequestURI(desktopFullUrl)
	if err := client.DoTimeout(req, res, 30*time.Second); err != nil {
		//log.Println(err)
		return nil,err
	}
	results = append(results,  string(res.Body()))

	return results, nil


}


type whoisHistory struct {
	Records []struct {
		AdministrativeContact struct {
			City         string `json:"city"`
			Country      string `json:"country"`
			Email        string `json:"email"`
			Fax          string `json:"fax"`
			FaxExt       string `json:"faxExt"`
			Name         string `json:"name"`
			Organization string `json:"organization"`
			PostalCode   string `json:"postalCode"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
			Street       string `json:"street"`
			Telephone    string `json:"telephone"`
			TelephoneExt string `json:"telephoneExt"`
		} `json:"administrativeContact"`
		Audit struct {
			CreatedDate string `json:"createdDate"`
			UpdatedDate string `json:"updatedDate"`
		} `json:"audit"`
		BillingContact struct {
			City         string `json:"city"`
			Country      string `json:"country"`
			Email        string `json:"email"`
			Fax          string `json:"fax"`
			FaxExt       string `json:"faxExt"`
			Name         string `json:"name"`
			Organization string `json:"organization"`
			PostalCode   string `json:"postalCode"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
			Street       string `json:"street"`
			Telephone    string `json:"telephone"`
			TelephoneExt string `json:"telephoneExt"`
		} `json:"billingContact"`
		CleanText          string   `json:"cleanText"`
		CreatedDateISO8601 string   `json:"createdDateISO8601"`
		CreatedDateRaw     string   `json:"createdDateRaw"`
		DomainName         string   `json:"domainName"`
		DomainType         string   `json:"domainType"`
		ExpiresDateISO8601 string   `json:"expiresDateISO8601"`
		ExpiresDateRaw     string   `json:"expiresDateRaw"`
		NameServers        []string `json:"nameServers"`
		RawText            string   `json:"rawText"`
		RegistrantContact  struct {
			City         string `json:"city"`
			Country      string `json:"country"`
			Email        string `json:"email"`
			Fax          string `json:"fax"`
			FaxExt       string `json:"faxExt"`
			Name         string `json:"name"`
			Organization string `json:"organization"`
			PostalCode   string `json:"postalCode"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
			Street       string `json:"street"`
			Telephone    string `json:"telephone"`
			TelephoneExt string `json:"telephoneExt"`
		} `json:"registrantContact"`
		RegistrarName    string   `json:"registrarName"`
		Status           []string `json:"status"`
		TechnicalContact struct {
			City         string `json:"city"`
			Country      string `json:"country"`
			Email        string `json:"email"`
			Fax          string `json:"fax"`
			FaxExt       string `json:"faxExt"`
			Name         string `json:"name"`
			Organization string `json:"organization"`
			PostalCode   string `json:"postalCode"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
			Street       string `json:"street"`
			Telephone    string `json:"telephone"`
			TelephoneExt string `json:"telephoneExt"`
		} `json:"technicalContact"`
		UpdatedDateISO8601 string `json:"updatedDateISO8601"`
		UpdatedDateRaw     string `json:"updatedDateRaw"`
		WhoisServer        string `json:"whoisServer"`
		ZoneContact        struct {
			City         string `json:"city"`
			Country      string `json:"country"`
			Email        string `json:"email"`
			Fax          string `json:"fax"`
			FaxExt       string `json:"faxExt"`
			Name         string `json:"name"`
			Organization string `json:"organization"`
			PostalCode   string `json:"postalCode"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
			Street       string `json:"street"`
			Telephone    string `json:"telephone"`
			TelephoneExt string `json:"telephoneExt"`
		} `json:"zoneContact"`
	} `json:"records"`
	RecordsCount int64 `json:"recordsCount"`
}



func GetWhoisHistory(domain string) (interface{} , error)  {


	endpoint := config.GetKey("WHOISXML_HISTORY_ENDPOINT")

	whoisHistoryPath := fmt.Sprintf("/api/v1?apiKey=%s&domainName=%s&mode=purchase&outPutFormat=json" , config.GetKey("WHOISXML_AUTH") , domain  )
	fullUrl := "https://"+endpoint + whoisHistoryPath

	whObj := new(whoisHistory)


	_, err := helpers.GetJson(fullUrl, whObj , "","")
	if err != nil {
		return nil, err
	}

	return whObj , nil

}



type Whois struct {
	WhoisRecord struct {
		AdministrativeContact struct {
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			Organization string `json:"organization"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
		} `json:"administrativeContact"`
		Audit struct {
			CreatedDate string `json:"createdDate"`
			UpdatedDate string `json:"updatedDate"`
		} `json:"audit"`
		ContactEmail          string `json:"contactEmail"`
		CreatedDate           string `json:"createdDate"`
		CreatedDateNormalized string `json:"createdDateNormalized"`
		CustomField1Name      string `json:"customField1Name"`
		CustomField1Value     string `json:"customField1Value"`
		CustomField2Name      string `json:"customField2Name"`
		CustomField2Value     string `json:"customField2Value"`
		CustomField3Name      string `json:"customField3Name"`
		CustomField3Value     string `json:"customField3Value"`
		DomainName            string `json:"domainName"`
		DomainNameExt         string `json:"domainNameExt"`
		EstimatedDomainAge    int64  `json:"estimatedDomainAge"`
		ExpiresDate           string `json:"expiresDate"`
		ExpiresDateNormalized string `json:"expiresDateNormalized"`
		Footer                string `json:"footer"`
		Header                string `json:"header"`
		NameServers           struct {
			HostNames []string      `json:"hostNames"`
			Ips       []interface{} `json:"ips"`
			RawText   string        `json:"rawText"`
		} `json:"nameServers"`
		ParseCode  int64  `json:"parseCode"`
		RawText    string `json:"rawText"`
		Registrant struct {
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			Organization string `json:"organization"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
		} `json:"registrant"`
		RegistrarIANAID string `json:"registrarIANAID"`
		RegistrarName   string `json:"registrarName"`
		RegistryData    struct {
			Audit struct {
				CreatedDate string `json:"createdDate"`
				UpdatedDate string `json:"updatedDate"`
			} `json:"audit"`
			CreatedDate           string `json:"createdDate"`
			CreatedDateNormalized string `json:"createdDateNormalized"`
			CustomField1Name      string `json:"customField1Name"`
			CustomField1Value     string `json:"customField1Value"`
			CustomField2Name      string `json:"customField2Name"`
			CustomField2Value     string `json:"customField2Value"`
			CustomField3Name      string `json:"customField3Name"`
			CustomField3Value     string `json:"customField3Value"`
			DomainName            string `json:"domainName"`
			ExpiresDate           string `json:"expiresDate"`
			ExpiresDateNormalized string `json:"expiresDateNormalized"`
			Footer                string `json:"footer"`
			Header                string `json:"header"`
			NameServers           struct {
				HostNames []string      `json:"hostNames"`
				Ips       []interface{} `json:"ips"`
				RawText   string        `json:"rawText"`
			} `json:"nameServers"`
			ParseCode             int64  `json:"parseCode"`
			RawText               string `json:"rawText"`
			RegistrarIANAID       string `json:"registrarIANAID"`
			RegistrarName         string `json:"registrarName"`
			Status                string `json:"status"`
			StrippedText          string `json:"strippedText"`
				UpdatedDate           string `json:"updatedDate"`
			UpdatedDateNormalized string `json:"updatedDateNormalized"`
			WhoisServer           string `json:"whoisServer"`
		} `json:"registryData"`
		Status           string `json:"status"`
		StrippedText     string `json:"strippedText"`
		TechnicalContact struct {
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			Organization string `json:"organization"`
			RawText      string `json:"rawText"`
			State        string `json:"state"`
		} `json:"technicalContact"`
		UpdatedDate           string `json:"updatedDate"`
		UpdatedDateNormalized string `json:"updatedDateNormalized"`
		WhoisServer           string `json:"whoisServer"`
	} `json:"WhoisRecord"`
}

func GetWhois(domain string) (interface{} , error)  {


	endpoint := config.GetKey("WHOISXML_GENERAL_ENDPOINT")

	whoisHistoryPath := fmt.Sprintf("/whoisserver/WhoisService?apiKey=%s&domainName=%s&outPutFormat=json" , config.GetKey("WHOISXML_AUTH") , domain  )
	fullUrl := "https://"+endpoint + whoisHistoryPath

	//log.Println(fullUrl)

	wObj := new(Whois)


	_, err := helpers.GetJson(fullUrl, wObj , "","")
	if err != nil {
		return nil, err
	}

	return wObj , nil


}


//
//type geoLocation struct {
//	As struct {
//		Asn    int64  `json:"asn"`
//		Domain string `json:"domain"`
//		Name   string `json:"name"`
//		Route  string `json:"route"`
//		Type   string `json:"type"`
//	} `json:"as"`
//	ConnectionType string   `json:"connectionType"`
//	Domains        []string `json:"domains"`
//	IP             string   `json:"ip"`
//	Isp            string   `json:"isp"`
//	Location       struct {
//		City       string  `json:"city"`
//		Country    string  `json:"country"`
//		GeonameID  int64   `json:"geonameId"`
//		Lat        float64 `json:"lat"`
//		Lng        float64 `json:"lng"`
//		PostalCode string  `json:"postalCode"`
//		Region     string  `json:"region"`
//		Timezone   string  `json:"timezone"`
//	} `json:"location"`
//}

//func GetIPGeolocation(ip string) (interface{} , error)  {
//
//	endpoint := config.GetKey("WHOISXML_GEOLOCATION_ENDPOINT")
//
//	whoisHistoryPath := fmt.Sprintf("/api/v1?apiKey=%s&ipAddress=%s&outPutFormat=json" , config.GetKey("WHOISXML_AUTH") , ip  )
//	fullUrl := "https://"+endpoint + whoisHistoryPath
//
//	gObj := new(geoLocation)
//
//	_, err := helpers.GetJson(fullUrl, gObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return gObj , nil
//
//
//}


//type dns struct {
//	DNSData struct {
//		Audit struct {
//			CreatedDate string `json:"createdDate"`
//			UpdatedDate string `json:"updatedDate"`
//		} `json:"audit"`
//		DNSRecords []struct {
//			AdditionalName string   `json:"additionalName"`
//			Address        string   `json:"address"`
//			Admin          string   `json:"admin"`
//			DNSType        string   `json:"dnsType"`
//			Expire         int64    `json:"expire"`
//			Flags          int64    `json:"flags"`
//			Host           string   `json:"host"`
//			Minimum        int64    `json:"minimum"`
//			Name           string   `json:"name"`
//			Priority       int64    `json:"priority"`
//			RRsetType      int64    `json:"rRsetType"`
//			RawText        string   `json:"rawText"`
//			Refresh        int64    `json:"refresh"`
//			Retry          int64    `json:"retry"`
//			Serial         int64    `json:"serial"`
//			Strings        []string `json:"strings"`
//			Tag            string   `json:"tag"`
//			Target         string   `json:"target"`
//			TTL            int64    `json:"ttl"`
//			Type           int64    `json:"type"`
//			Value          string   `json:"value"`
//		} `json:"dnsRecords"`
//		DNSTypes   string  `json:"dnsTypes"`
//		DomainName string  `json:"domainName"`
//		Types      []int64 `json:"types"`
//	} `json:"DNSData"`
//}

//func GetDNS(domain string) (interface{} , error)  {
//
//
//	endpoint := config.GetKey("WHOISXML_GENERAL_ENDPOINT")
//
//	whoisHistoryPath := fmt.Sprintf("/whoisserver/DNSService?apiKey=%s&domainName=%s&type=_all&outPutFormat=json" , config.GetKey("WHOISXML_AUTH") , domain  )
//	fullUrl := "https://"+endpoint + whoisHistoryPath
//
//	dObj := new(dns)
//
//
//	_, err := helpers.GetJson(fullUrl, dObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return dObj , nil
//
//}


