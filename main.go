package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"bufio"
	"net/http"
	"encoding/base64"
	"sync"
	"crypto/tls"
)

type structs struct {
	Time float64
	Dcfd  string
	Sdcfd string
	Xprops string
	Xconst string
}


func joiner(token string, invite string) {
	payload := map[string]string{}
	xp,_ := json.Marshal(payload)
	Cookie := Build_cookie()
	Cookies := "__dcfduid=" + Cookie.Dcfd + "; " + "__sdcfduid=" + Cookie.Sdcfd + "; "
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/invites/"+invite+"", bytes.NewBuffer(xp))
	if err != nil {
		log.Fatal(err)
	}
	for x,o := range map[string]string{
		"accept": "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en-NL;q=0.9,en-GB;q=0.8",
		"authorization": token,
		"content-length": "2",
		"content-type": "application/json",
		"cookie": Cookies,
		"origin": "https://discord.com",
		"referer": "https://discord.com/channels/@me/",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
		"x-context-properties": Build_Xheader().Xconst,
		"x-debug-options": "bugReporterEnabled",
		"x-discord-locale": "en-US",
		"x-super-properties": Build_Xheader().Xprops,
	} {
		req.Header.Set(x,o)
	}
	resp, err := Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("\033[32m┃"+r+" (\033[32m+\033[39m) Joined Server "+c+"| "+r+"", token[:40], " "+c+"|"+r+" gg"+c+"/"+r+""+invite+"")
	} else if resp.StatusCode == 429 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var ResponseBody structs
		json.Unmarshal(body, &ResponseBody)
		timeout := ResponseBody.Time
		fmt.Println("\033[33m┃ (\033[33m/\033[39m) Rate Limmit "+c+"|"+r+" ["+c+"TIME"+r+"]: ", timeout)
	} else {
		fmt.Println("\033[31m┃ "+r+"(\033[31mx"+r+") Somthing Whent Wrong "+c+"|"+r+" ", resp[:100], " | Captcha")
	}



}



func  Build_Xheader() structs {
	Xheader := structs{}
	xconststr := `{"location":"Invite Button Embed","location_guild_id":null,"location_channel_id":"","location_channel_type":3,"location_message_id":""}`
	xpropsstr := `{"os":"Windows","browser":"Discord Client","release_channel":"stable","client_version":"1.0.9006","os_version":"10.0.22000","os_arch":"x64","system_locale":"en-US","client_build_number":151638,"client_event_source":null}`
	Xheader.Xconst = base64.StdEncoding.EncodeToString([]byte(xconststr))
	Xheader.Xprops = base64.StdEncoding.EncodeToString([]byte(xpropsstr))
	return Xheader
}


func Build_cookie() structs {
	req, err := http.Get("https://discord.com")
	if err != nil {
		log.Fatal(err)
		CookieNil := structs{}
		return CookieNil
	}
	defer req.Body.Close()

	Cookie := structs{}
	if req.Cookies() != nil {
		for _, cookie := range req.Cookies() {
			if cookie.Name == "__dcfduid" {
				Cookie.Dcfd = cookie.Value
			}
			if cookie.Name == "__sdcfduid" {
				Cookie.Sdcfd = cookie.Value
			}
		}
	}
	return Cookie
}


func read_tokens() ([]string, error) {
	file, err := os.Open("tokens.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls") 
	cmd.Stdout = os.Stdout
	cmd.Run()
}




var c = "\033[36m"
var r = "\033[39m"
var Client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS13,
		},
		//Proxy: http.ProxyURL(p),
	},
}



func main() {
	cls()
	logo := ``+c+`
┃ _______________     ___________________________   __________________ 
┃ __  ____/_  __ \    ______  /_  __ \___  _/__  | / /__  ____/__  __ \
┃ _  / __ _  / / /    ___ _  /_  / / /__  / __   |/ /__  __/  __  /_/ /
┃ / /_/ / / /_/ /     / /_/ / / /_/ /__/ /  _  /|  / _  /___  _  _, _/ 
┃ \____/  \____/      \____/  \____/ /___/  /_/ |_/  /_____/  /_/ |_|  
┃	
┃  `+r+`[`+c+`https://github.com/yaboipy/go-joiner`+r+`]`
	fmt.Print(logo)
	scn := bufio.NewScanner(os.Stdin)
	lines, err := read_tokens()
	if err != nil {
		log.Fatal(err)
	}	
	fmt.Print("\n"+c+"┃\n"+c+"┃"+r+"	("+c+"-"+r+")  discord.gg"+c+"/"+r+"")
	scn.Scan()
	invite := scn.Text()
	var wg sync.WaitGroup
	wg.Add(len(lines))
	for i := 0; i < len(lines); i++ {
		joiner(lines[i], invite)
	}
}
