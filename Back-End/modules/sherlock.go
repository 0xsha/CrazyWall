package modules

import (
	"back-end/config"
	"bytes"
	//"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type AccountData struct {
	Output string `json:"output"`
}


func GetSocialAccounts(username string) (interface{} , error) {

	endpoint := config.GetKey("SHERLOCK_ENDPOINT")

	sherlock :=  fmt.Sprintf(":8000/cli/")

	fullUrl := "http://"+endpoint+sherlock

	//log.Println(fullUrl)

	//sObj := new(AccountData)

	params := ` --site reddit --site hackster  --site tryhackme --site pr0gramm --site leasehackr --site 9GAG --site About.me --site Academia.edu --site Asciinema --site Archive.org --site AskFM --site BitBucket --site BitCoinForum --site Blogger --site BodyBuilding --site BuzzFeed --site Carbonmade --site CloudflareCommunity --site Codepen --site Cracked --site 'DEV Community' --site DailyMotion --site 'Docker Hub' --site Dribbble --site Etsy --site Flickr --site Football --site Freelancer --site GitHub --site GitLab --site Gitee --site GoodReads --site Gravatar --site Gumroad --site HackTheBox --site Hackaday --site HackerOne --site HackerNews --site ICQ --site Imgur --site Keybase --site LeetCode --site LiveJournal --site Medium --site PCGamer --site Pastebin --site Patreon --site PlayStore --site ProductHunt --site ResearchGate --site Signal --site SlideShare --site Spotify --site steam --site Telegram --site Tinder --site TradingView --site Trello --site Twitch --site VK --site Venmo --site VirusTotal --site Wix`
	args := username + " " + params


	var jsonStr = []byte(`{"args" : "`+args +`"}`)




	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonStr))
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
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))


	//if err := json.Unmarshal(body, sObj); err != nil {
	//	return nil,err
	//}


	return string(body) , nil



}
