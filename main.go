package main


import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	"os"
	"os/exec"
	"sync"
	"time"
)


type structs struct {
	Raider struct {
		Interval int `json:"interval"`
		Hex 	 string `json:"hex"`
	} `json:"Raider"`
	Joiner struct {
		Proxy 	 string `json:"proxy"`
	} `json:"Joiner"`
	
	Time    float64
	Dcfd 	string
	Sdcfd 	string
	Xconst  string
	Xprops  string
	
}

var (
	c = "\033[36m"
	r = "\033[39m"
	proxy = config().Joiner.Proxy
	Cookies = "__dcfduid=" + Build_cookie().Dcfd + "; " + "__sdcfduid=" + Build_cookie().Sdcfd + "; "
)


func joiner(token string, invite string) {
	//p, _ := url.Parse("http://" + proxy)
	Client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS13,
			},
			//Proxy: http.ProxyURL(p),
		},
	}
	payload := map[string]string{}
	xp,_ := json.Marshal(payload)
	Cookie := Build_cookie()
	Cookies := "__dcfduid=" + Cookie.Dcfd + "; " + "__sdcfduid=" + Cookie.Sdcfd + "; "
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/invites/"+invite+"", bytes.NewBuffer(xp))
	cerr(err)
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
	cerr(err)
	if resp.StatusCode == 200 {
		fmt.Println("\033[32m┃"+r+" (\033[32m+\033[39m) Joined Server "+c+"| "+r+"", token[:40], " "+c+"|"+r+" gg"+c+"/"+r+""+invite+"")
	} else if resp.StatusCode == 429 {
		body, err := ioutil.ReadAll(resp.Body)
		cerr(err)
		var ResponseBody structs
		json.Unmarshal(body, &ResponseBody)
		timeout := ResponseBody.Time
		fmt.Println("\033[33m┃ (\033[33m/\033[39m) Rate Limmit "+c+"|"+r+" ["+c+"TIME"+r+"]: ", timeout)
	} else {
		fmt.Println("\033[31m┃ "+r+"(\033[31mx"+r+") Somthing Whent Wrong "+c+"|"+r+" Captcha")
	}

}




func raider(token string, channel string, message string) {
	Client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS13,
			},
		},
	}
	hex := config().Raider.Hex
	hx,_ := Hex(hex)
	payload := map[string]string{
		"content": message + " | " + hx,
	}
	xp,_ := json.Marshal(payload)
	interval := config().Raider.Interval
	for true {
		time.Sleep(time.Duration(interval))
		req, err := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+channel+"/messages", bytes.NewBuffer(xp))
		cerr(err)
		for x,o := range map[string]string{
			"accept": "*/*",
			"accept-encoding": "gzip, deflate, br",
			"accept-language": "en-US,en-NL;q=0.9,en-GB;q=0.8",
			"authorization": token,
			"content-type": "application/json",
			"cookie": Cookies,
			"origin": "https://discord.com",
			"referer": "https://discord.com/channels/@me/"+channel+"",
			"sec-fetch-dest": "empty",
			"sec-fetch-mode": "cors",
			"sec-fetch-site": "same-origin",
			"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
			"x-debug-options": "bugReporterEnabled",
			"x-discord-locale": "en-US",
			"x-super-properties": Build_Xheader().Xprops,
		} {
			req.Header.Set(x,o)
		}
		resp, err := Client.Do(req)
		cerr(err)
		if resp.StatusCode == 200 {
			fmt.Println("("+c+"+"+r+") Sent Message | ", message)
		} else if resp.StatusCode == 429 {
			fmt.Println("RateLimit")
		} else {
			fmt.Println("Failed To Send")
		}
	}
}





func Build_cookie() structs {
	Client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MaxVersion: tls.VersionTLS13,
			},
		},
	}
	req, err := http.NewRequest("GET", "https://discord.com", nil)
	if err != nil {
		cerr(err)
		CookieNil := structs{}
		return CookieNil
	}
	for x,o := range map[string]string{
		"accept":" */*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en;q=0.9",
		"content-type": "application/json",
		"sec-ch-ua": `Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105`,
		"sec-ch-ua-mobile": "?0",
		"sec-ch-ua-platform": "Windows",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
	} {
		req.Header.Set(x,o)
	}
	resp, err := Client.Do(req)
	cerr(err)
	defer resp.Body.Close()
	Cookie := structs{}
	if resp.Cookies() != nil {
		for _, cookie := range resp.Cookies() {
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


func Hex(x string) (string, error) {
	bytes := []byte(x)
	if _, err := rand.Read(bytes); err != nil {
	  return "", err
	}
	return hex.EncodeToString(bytes), nil
}


  func  Build_Xheader() structs {
	Xheader := structs{}
	xconststr := `{"location":"Invite Button Embed","location_guild_id":null,"location_channel_id":"","location_channel_type":3,"location_message_id":""}`
	xpropsstr := `{"os":"Windows","browser":"Discord Client","release_channel":"stable","client_version":"1.0.9006","os_version":"10.0.22000","os_arch":"x64","system_locale":"en-US","client_build_number":151638,"client_event_source":null}`
	Xheader.Xconst = base64.StdEncoding.EncodeToString([]byte(xconststr))
	Xheader.Xprops = base64.StdEncoding.EncodeToString([]byte(xpropsstr))
	return Xheader
}


func config() structs {
	var config structs
	conf, err := os.Open("config.json")
	defer conf.Close()
	cerr(err)
	xp := json.NewDecoder(conf)
	xp.Decode(&config)
	return config

}



func read_tokens() ([]string, error) {
	file, err := os.Open("tokens.txt")
	cerr(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


func cerr(err error) {
	if err != nil {
		log.Fatal(err)
	}	
}



func cls() {
	cmd := exec.Command("cmd", "/c", "cls") 
	cmd.Stdout = os.Stdout
	cmd.Run()
}



func main() {
	cls()
	logo := `
		____`+c+`_____`+r+`___`+c+`___`+r+`_____`+c+`___`+r+`____`+c+`____`+r+`____`+c+`___`+r+`____`+c+`___ `+r+`
		__  `+c+`____/`+r+`_`+c+`  __ \`+r+`__`+c+`  __ \`+r+`__`+c+`    |`+r+`___`+c+`  _/`+r+`__`+c+`  __ \`+r+`
		_ `+c+` / __ `+r+`_  `+c+`/ / /`+r+`_  `+c+`/_/ /`+r+`_ `+c+` /| |`+r+`__`+c+`  / `+r+`__`+c+`  / / /
		/ /_/ / / /_/ /`+r+`_  `+c+`_, _/`+r+`_  `+c+`___ |_/ /  `+r+`_  `+c+`/_/ / 
		\____/  \____/ /_/ |_| /_/  |_/___/  /_____/  

				`+r+`[`+c+`1`+r+`] Joiner
				`+r+`[`+c+`2`+r+`] Raider 
	
	`
	fmt.Println(logo)
	scn := bufio.NewScanner(os.Stdin)
	lines, err := read_tokens()
	cerr(err)
	cerr(err)
	fmt.Print("	["+c+">"+r+"] Choice: ")
	scn.Scan()
	choice := scn.Text()
	if choice == "1" {
		fmt.Print("	["+c+">"+r+"] discord"+c+"/"+r+"")
		scn.Scan()
		inv := scn.Text()
		var wg sync.WaitGroup
		wg.Add(len(lines))
		for i := 0; i < len(lines); i++ {
			go func(i int) {
				joiner(lines[i], inv)
			}(i)
		}
		select{}
	} else if choice == "2" {
		fmt.Print("	["+c+">"+r+"] Channel ID: ")
		scn.Scan()
		chnl := scn.Text()
		fmt.Print("	["+c+">"+r+"] Message: ")
		scn.Scan()
		msg := scn.Text()
		for i := 0; i < len(lines); i++ {
			go func(i int) {
				for {
					raider(lines[i], chnl, msg)
				}
			}(i)
		}
		select{}
	} else  {
		fmt.Println("Wrong Input")
		time.Sleep(1 *time.Second)
		main()
	}
}
