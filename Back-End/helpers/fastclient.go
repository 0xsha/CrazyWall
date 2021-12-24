package helpers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"time"
)

func GetJson(url string, obj interface{} , authType string , authCred string) (interface{} , error)  {

	client := fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	// different APIs different authorization
	if authType == "dehashed"{

		// for dehashed
		authKey := base64.StdEncoding.EncodeToString([]byte(authCred))
		req.Header.Add("Authorization","Basic "+authKey)
		req.Header.Add("Accept", "application/json")

	}

	if authType == "wappalyzer"{

		// for wappalyzer
		req.Header.Add("x-api-key", authCred)
		req.Header.Add("Accept", "application/json")

	}


	req.SetRequestURI(url)
	if err := client.DoTimeout(req, res, 30*time.Second); err != nil {
		//log.Println(err)
		return nil,err
	}




	//log.Println(string(res.Body()))

	if err := json.Unmarshal(res.Body(), obj); err != nil {
		return nil,err
	}

	return obj,nil
}
