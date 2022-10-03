package main

//----------------\
// Made by YABOI   \
// GO Token Joiner |
//----------------/


import (
	"bytes"
	"log"
	"fmt"
	"net/http"
	"encoding/base64"
	"encoding/json"
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

func joiner() {
	Client := cycletls.Init()
	cookie := "__dcfduid=98fda5d03e2e11ed889365ed1a91b671; __sdcfduid=dcfduid=98fda5d03e2e11ed889365ed1a91b6715714520f372b42d532f848947adbb3f21ec8288fe425fc3cede4e91af51c065e; __cfruid=4dc76b3ea4b723672629a2cbdb7b3a3702265855-1664783924"
	xconstr := `{"location":"Join Guild","location_guild_id":"`+guild+`","location_channel_id":"`+channel+`","location_channel_type":0}`
	xstring := `{"os":"Windows","browser":"Discord Client","release_channel":"stable","client_version":"1.0.9006","os_version":"10.0.22000","os_arch":"x64","system_locale":"en-US","client_build_number":150347,"client_event_source":null}`
	agent := "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36"
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/invites/"+invite+"?inputValue="+invite+"&with_counts=true&with_expiration=true", nil)
	if err != nil {
		log.Fatal(err)
	}
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-GB,en-US;q=0.9",
		"authorization": token,
		"cookie": cookie,
		"referer": "https://discord.com/channels/@me/",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": agent,
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": base64.StdEncoding.EncodeToString([]byte(xstring)),
	} {
		req.Header.Set(x,o)
	}
	resp, err := Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Response.StatusCode == 200 {
		fmt.Println("[Xprops]> " + base64.StdEncoding.EncodeToString([]byte(xstring)))
		payload := map[string]string{}
		xp,_ := json.Marshal(payload) 
		req, err := http.NewRequest("POST", "https://discord.com/api/v9/invites/"+invite+"", bytes.NewBuffer(xp))
		if err != nil {
			log.Fatal(err)
		}
		for x,o := range map[string]string{
			"accept": "*/*",
            "accept-encoding": "gzip, deflate, br",
            "accept-language": "en-GB,en-US;q=0.9",
            "authorization": token,
            "content-length": "2",
            "content-type": "application/json",
            "cookie": cookie,
            "origin": "https://discord.com",
            "referer": "https://discord.com/channels/@me/",
            "sec-fetch-dest": "empty",
            "sec-fetch-mode": "cors",
            "sec-fetch-site": "same-origin",
            "user-agent": agent,
            "x-context-properties": base64.StdEncoding.EncodeToString([]byte(xconstr)),
            "x-debug-options": "bugReporterEnabled",
            "x-discord-locale": "en-US",
            "x-super-properties": base64.StdEncoding.EncodeToString([]byte(xstring)),
		} {
			req.Header.Set(x,o)
		}
		resp, err := Client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 200 {
			fmt.Println("[>] Joined | " + token)
		} else {
			fmt.Println("[>] Failed | " + token + " | " + resp)
		}

	}
}



var client = cycletls.Init()

func main() {
	for token := range token {
		joiner()
	}
}
