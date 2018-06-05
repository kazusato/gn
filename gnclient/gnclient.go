/**
 * Copyright 2018 Kazuhiko Sato
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

 package gnclient

import (
	"net/url"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"net/http/cookiejar"
	"log"
	"strings"
)

type GnClient struct {
	Url, UserId, Password string
	jar *cookiejar.Jar
	client *http.Client
}

const debug = false
const loginUrl = "index.php"
const cmdUrl = "cmd.php"
const itinJaUrl = "pnr_itinerary_ja.php"
const itinEnUrl = "pnr_itinerary_en.php"

func NewClientFromConfig(config *Config) *GnClient {
	return NewClient(config.Url, config.UserId, config.Password)
}

func NewClient(Url string, UserId string, Password string) *GnClient {
	gn := GnClient{Url: Url, UserId: UserId, Password: Password}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	gn.jar = jar
	gn.client = &http.Client{Jar: jar}

	return &gn
}

func (gn GnClient) Connect() error {
	data := url.Values{}
	data.Set("default_command", "login")
	data.Set("login_userid", gn.UserId)
	data.Set("login_password", gn.Password)
	data.Set("command_login.x", "5")
	data.Set("command_login.y", "18")

	resp, err := gn.client.PostForm(gn.Url + loginUrl, data)
	if err != nil {
		log.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if debug {
		log.Println(string(body))
	}

	defer resp.Body.Close()

	return nil
}

func (gn GnClient) SendCommand(cmd string) (string, error) {
	data := defaultForm()
	data.Set("cmdrequest", cmd)

	resp, err := gn.client.PostForm(gn.Url + cmdUrl, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	respStr := string(body)
	decodedStr := decodeCmdRes(respStr)

	defer resp.Body.Close()

	return decodedStr, nil
}

func decodeCmdRes(str string) string {
	token := strings.Split(str, "&")
	for _, value := range token {
		keyVal := strings.Split(value, "=")
		paramKey := keyVal[0]
		paramValue := keyVal[1]

		if paramKey == "cmdres" {
			decoded, err := url.QueryUnescape(paramValue)
			if err != nil {
				log.Println(err)
			}
			return strings.Replace(decoded, "\r", "\r\n", -1)
		}
	}

	return ""
}

func defaultForm() url.Values {
	data := url.Values{}
	data.Set("topck", "1")
	data.Set("end", "1")
	data.Set("cmdhistory01", "0asa")
	data.Set("cmdhistory02", "0asa")
	data.Set("cmdhistory03", "0asa")
	data.Set("cmdhistory04", "0asa")
	data.Set("cmdhistory05", "0asa")
	data.Set("cmdhistory06", "0asa")
	data.Set("cmdhistory07", "0asa")
	data.Set("cmdhistory09", "0asa")
	data.Set("cmdhistory10", "0asa")
	data.Set("cmdhistory11", "0asa")
	data.Set("cmdhistory12", "0asa")
	data.Set("cmdhistory13", "0asa")
	data.Set("cmdhistory14", "0asa")
	data.Set("cmdhistory15", "0asa")
	data.Set("cmdhistory16", "0asa")
	data.Set("cmdhistory17", "0asa")
	data.Set("cmdhistory18", "0asa")
	data.Set("cmdhistory19", "0asa")
	data.Set("cmdhistory20", "0asa")
	data.Set("memflg", "0")
	data.Set("command", "flash_submit")
	data.Set("cmderror", "0")
	data.Set("cmdtab1", "0")
	data.Set("cmdtab2", "0")
	data.Set("cmdtab3", "0")
	data.Set("cmdtab4", "0")
	data.Set("cmdtab5", "0")
	data.Set("cmdtab6", "0")
	data.Set("cmdtab7", "0")
	data.Set("cmdtab8", "0")
	data.Set("cmdtab9", "0")
	data.Set("cmdtab10", "0")
	data.Set("cmdtab11", "0")
	data.Set("cmdtab12", "0")
	data.Set("cmdtab13", "0")
	data.Set("cmdtab14", "0")
	data.Set("cmdtab15", "0")
	data.Set("cmdtab16", "0")
	data.Set("cmdtab17", "0")
	data.Set("cmdtab18", "0")
	data.Set("cmdtab19", "0")
	data.Set("cmdtab20", "0")
	data.Set("cmdtab21", "0")
	data.Set("cmdtab22", "0")
	data.Set("cmdtab23", "0")
	data.Set("cmdtab24", "0")
	data.Set("cmdtab25", "0")
	data.Set("cmdtab26", "0")
	data.Set("cmdtab27", "0")
	data.Set("cmdtab28", "0")
	data.Set("cmdtab29", "0")
	data.Set("cmdtab30", "0")
	data.Set("cmdres", "` ")
	data.Set("cmdrequest", "")

	return data
}
