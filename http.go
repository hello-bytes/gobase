package gobase

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(url string) ([]byte, int) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("do client err->%v\n", err)
		return nil, 0
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("do client err->%v\n", err)
		return nil, 0
	}

	return body, 200
}

func HttpPost(urlPath string, params *map[string]string) ([]byte, int) {
	strparam := ""
	for key, val := range *params {
		urlKey := url.QueryEscape(key)
		urlVal := url.QueryEscape(val)
		if len(strparam) == 0 {
			strparam = urlKey + "=" + urlVal
		} else {
			strparam = strparam + "&" + urlKey + "=" + urlVal
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(strparam))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("do client err->%v\n", err)
		return nil, 1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("do client err->%v\n", err)
		return nil, 2
	}

	return body, 200
}
