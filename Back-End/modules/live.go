package modules

import (
	"back-end/config"
	"back-end/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)


type NewsApiResp struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		//Content     string    `json:"content"`
	} `json:"articles"`
}



func GetNews(keyword string) (interface{} , error)  {

	endpoint := config.GetKey("NEWSAPI_ENDPOINT")
	fbPath :=  fmt.Sprintf(`/v2/everything?qInTitle="%s"&sortBy=publishedAt&pageSize=100&language=en&apiKey=%s` , strings.Replace(keyword, " ", "%20", -1) , config.GetKey("NEWSAPI_AUTH"))


	//fbPath :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=facebook" , keyword, config.GetKey("SOCIAL_SEARCHER_AUTH") , limit)

	fullUrl := "https://"+endpoint+fbPath

	//log.Println(fullUrl)

	nObj := new(NewsApiResp)


	_, err := helpers.GetJson(fullUrl, nObj , "","")
	if err != nil {
		return nil, err
	}

	return nObj , nil


}


type SocialData struct {
	Meta struct {
		HTTPCode  int64  `json:"http_code"`
		Limit     int64  `json:"limit"`
		Network   string `json:"network"`
		Page      int64  `json:"page"`
		QueryType string `json:"query_type"`
		Requestid string `json:"requestid"`
	} `json:"meta"`
	Posts []struct {
		Image      string `json:"image"`
		Lang       string `json:"lang"`
		Network    string `json:"network"`
		Popularity []struct {
			Count int64  `json:"count"`
			Name  string `json:"name"`
		} `json:"popularity"`
		Posted    string `json:"posted"`
		Postid    string `json:"postid"`
		Sentiment string `json:"sentiment"`
		Tags      []struct {
			Text string `json:"text"`
			URL  string `json:"url"`
		} `json:"tags"`
		Text string `json:"text"`
		Type string `json:"type"`
		URL  string `json:"url"`
		Urls []struct {
			ShortURL string `json:"short_url"`
			Text     string `json:"text"`
			URL      string `json:"url"`
		} `json:"urls"`
		User struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Influence   struct {
				Count int64  `json:"count"`
				Name  string `json:"name"`
			} `json:"influence"`
			Location string `json:"location"`
			Name     string `json:"name"`
			URL      string `json:"url"`
			Userid   string `json:"userid"`
		} `json:"user"`
		UserMentions []struct {
			Text string `json:"text"`
			URL  string `json:"url"`
		} `json:"user_mentions"`
	} `json:"posts"`
}

type FullSocialData struct {

	FullData  []SocialData
}

func GetSocial(keyword string) ([]FullSocialData , error) {


	endpoint := config.GetKey("SOCIAL_SEARCHER_ENDPOINT")

	firstUrl :=  fmt.Sprintf("/v2/search?q='%s'&key=%s&limit=%d&network=instagram,youtube,twitter,reddit,facebook" ,  strings.Replace(keyword, " ", "%20", -1) , config.GetKey("SOCIAL_SEARCHER_AUTH") , 100)

	fullUrl := "https://"+endpoint+firstUrl
	sObj := new(SocialData)


	req, err := http.NewRequest("GET", fullUrl , nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)



	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	err = json.NewDecoder(resp.Body).Decode(sObj)
	if err != nil {
		//log.Println(err)
		return nil,err

	}


	fullData := make([]FullSocialData,4)

	requestID := sObj.Meta.Requestid

	fullData[0].FullData = append(fullData[0].FullData , *sObj)




	for i:=1 ; i<4; i++ {



		firstUrl :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=youtube,instagram,twitter,reddit,facebook&page=%d&requestid=%s" ,  strings.Replace(keyword, " ", "%20", -1) , config.GetKey("SOCIAL_SEARCHER_AUTH") , 100 , i ,requestID )
		fullUrl := "https://"+endpoint+firstUrl

		sObj := new(SocialData)

		req, err := http.NewRequest("GET", fullUrl , nil)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)



		err = json.NewDecoder(resp.Body).Decode(sObj)
		if err != nil {

			//log.Println(err)
			return nil,err

		}

		fullData[i].FullData = append(fullData[i].FullData , *sObj)



	}

	//log.Println(fullData[1])


	return fullData , nil

}



//type Facebook struct {
//	//Meta struct {
//	//	HTTPCode  int64  `json:"http_code"`
//	//	Limit     int64  `json:"limit"`
//	//	Network   string `json:"network"`
//	//	Page      int64  `json:"page"`
//	//	QueryType string `json:"query_type"`
//	//	Requestid string `json:"requestid"`
//	//} `json:"meta"`
//	Posts []struct {
//		Lang      string `json:"lang"`
//		Network   string `json:"network"`
//		Posted    string `json:"posted"`
//		Postid    string `json:"postid"`
//		Sentiment string `json:"sentiment"`
//		Text      string `json:"text"`
//		Type      string `json:"type"`
//		URL       string `json:"url"`
//		Urls      []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"urls"`
//		User struct {
//			Name string `json:"name"`
//			URL  string `json:"url"`
//		} `json:"user"`
//	} `json:"posts"`
//}
//func GetFacebook(keyword string , limit int) (interface{} , error)  {
//
//
//	endpoint := config.GetKey("SOCIAL_SEARCHER_ENDPOINT")
//
//	fbPath :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=facebook" ,  strings.Replace(keyword, " ", "%20", -1) , config.GetKey("SOCIAL_SEARCHER_AUTH") , limit)
//
//	fullUrl := "https://"+endpoint+fbPath
//	fObj := new(Facebook)
//
//	_, err := helpers.GetJson(fullUrl, fObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return fObj , nil
//
//}
//
//
//type reddit struct {
//	Meta struct {
//		HTTPCode  int64  `json:"http_code"`
//		Limit     int64  `json:"limit"`
//		Network   string `json:"network"`
//		Page      int64  `json:"page"`
//		QueryType string `json:"query_type"`
//		Requestid string `json:"requestid"`
//	} `json:"meta"`
//	Posts []struct {
//		Image      string `json:"image"`
//		Lang       string `json:"lang"`
//		Network    string `json:"network"`
//		Popularity []struct {
//			Count int64  `json:"count"`
//			Name  string `json:"name"`
//		} `json:"popularity"`
//		Posted    string `json:"posted"`
//		Postid    string `json:"postid"`
//		Sentiment string `json:"sentiment"`
//		Text      string `json:"text"`
//		Type      string `json:"type"`
//		URL       string `json:"url"`
//		Urls      []struct {
//			Caption string `json:"caption"`
//			Text    string `json:"text"`
//			URL     string `json:"url"`
//		} `json:"urls"`
//		User struct {
//			Name   string `json:"name"`
//			URL    string `json:"url"`
//			Userid string `json:"userid"`
//		} `json:"user"`
//	} `json:"posts"`
//}
//
//
//func GetReddit(keyword string , limit int) (interface{} , error)  {
//
//
//	endpoint := config.GetKey("SOCIAL_SEARCHER_ENDPOINT")
//
//	redditPath :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=reddit" , strings.Replace(keyword, " ", "%20", -1), config.GetKey("SOCIAL_SEARCHER_AUTH") , limit)
//
//	fullUrl := "https://"+endpoint+redditPath
//	rObj := new(reddit)
//
//	_, err := helpers.GetJson(fullUrl, rObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return rObj , nil
//
//}
//
//
//
//type youtube struct {
//	Meta struct {
//		HTTPCode  int64  `json:"http_code"`
//		Limit     int64  `json:"limit"`
//		Network   string `json:"network"`
//		Page      int64  `json:"page"`
//		QueryType string `json:"query_type"`
//		Requestid string `json:"requestid"`
//	} `json:"meta"`
//	Posts []struct {
//		Image     string `json:"image"`
//		Lang      string `json:"lang"`
//		Network   string `json:"network"`
//		Posted    string `json:"posted"`
//		Postid    string `json:"postid"`
//		Sentiment string `json:"sentiment"`
//		Tags      []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"tags"`
//		Text string `json:"text"`
//		Type string `json:"type"`
//		URL  string `json:"url"`
//		Urls []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"urls"`
//		User struct {
//			Name   string `json:"name"`
//			URL    string `json:"url"`
//			Userid string `json:"userid"`
//		} `json:"user"`
//	} `json:"posts"`
//}
//
//
//func GetYoutube(keyword string , limit int) (interface{} , error)  {
//
//
//	endpoint := config.GetKey("SOCIAL_SEARCHER_ENDPOINT")
//
//	ytPath :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=youtube" , strings.Replace(keyword, " ", "%20", -1), config.GetKey("SOCIAL_SEARCHER_AUTH") , limit)
//
//	fullUrl := "https://"+endpoint+ytPath
//	yObj := new(youtube)
//
//	_, err := helpers.GetJson(fullUrl, yObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return yObj , nil
//
//}
//
//
//type instagram struct {
//	Meta struct {
//		HTTPCode  int64  `json:"http_code"`
//		Limit     int64  `json:"limit"`
//		Network   string `json:"network"`
//		Page      int64  `json:"page"`
//		QueryType string `json:"query_type"`
//		Requestid string `json:"requestid"`
//	} `json:"meta"`
//	Posts []struct {
//		Image      string `json:"image"`
//		Lang       string `json:"lang"`
//		Network    string `json:"network"`
//		Popularity []struct {
//			Count int64  `json:"count"`
//			Name  string `json:"name"`
//		} `json:"popularity"`
//		Posted    string `json:"posted"`
//		Postid    string `json:"postid"`
//		Sentiment string `json:"sentiment"`
//		Tags      []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"tags"`
//		Text string `json:"text"`
//		Type string `json:"type"`
//		URL  string `json:"url"`
//		Urls []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"urls"`
//		User struct {
//			Name   string `json:"name"`
//			Userid string `json:"userid"`
//		} `json:"user"`
//		UserMentions []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"user_mentions"`
//	} `json:"posts"`
//}
//
//func GetInstagram(keyword string , limit int) (interface{} , error)  {
//
//
//	endpoint := config.GetKey("SOCIAL_SEARCHER_ENDPOINT")
//
//	igPath :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=instagram" , strings.Replace(keyword, " ", "%20", -1), config.GetKey("SOCIAL_SEARCHER_AUTH") , limit)
//
//	fullUrl := "https://"+endpoint+igPath
//	iObj := new(instagram)
//
//	_, err := helpers.GetJson(fullUrl, iObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return iObj , nil
//
//}
//
//
//
//type twitter struct {
//	Meta struct {
//		HTTPCode  int64  `json:"http_code"`
//		Limit     int64  `json:"limit"`
//		Network   string `json:"network"`
//		Page      int64  `json:"page"`
//		QueryType string `json:"query_type"`
//		Requestid string `json:"requestid"`
//	} `json:"meta"`
//	Posts []struct {
//		Image      string `json:"image"`
//		Lang       string `json:"lang"`
//		Network    string `json:"network"`
//		Popularity []struct {
//			Count int64  `json:"count"`
//			Name  string `json:"name"`
//		} `json:"popularity"`
//		Posted    string `json:"posted"`
//		Postid    string `json:"postid"`
//		Sentiment string `json:"sentiment"`
//		Tags      []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"tags"`
//		Text string `json:"text"`
//		Type string `json:"type"`
//		URL  string `json:"url"`
//		Urls []struct {
//			ShortURL string `json:"short_url"`
//			Text     string `json:"text"`
//			URL      string `json:"url"`
//		} `json:"urls"`
//		User struct {
//			Description string `json:"description"`
//			Image       string `json:"image"`
//			Influence   struct {
//				Count int64  `json:"count"`
//				Name  string `json:"name"`
//			} `json:"influence"`
//			Location string `json:"location"`
//			Name     string `json:"name"`
//			URL      string `json:"url"`
//			Userid   string `json:"userid"`
//		} `json:"user"`
//		UserMentions []struct {
//			Text string `json:"text"`
//			URL  string `json:"url"`
//		} `json:"user_mentions"`
//	} `json:"posts"`
//}
//
//func GetTwitter(keyword string , limit int) (interface{} , error)  {
//
//
//	endpoint := config.GetKey("SOCIAL_SEARCHER_ENDPOINT")
//
//	twitterPath :=  fmt.Sprintf("/v2/search?q=%s&key=%s&limit=%d&network=twitter" , strings.Replace(keyword, " ", "%20", -1), config.GetKey("SOCIAL_SEARCHER_AUTH") , limit)
//
//	fullUrl := "https://"+endpoint+twitterPath
//	tObj := new(twitter)
//
//	_, err := helpers.GetJson(fullUrl, tObj , "","")
//	if err != nil {
//		return nil, err
//	}
//
//	return tObj , nil
//
//}
//
