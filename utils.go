package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const UA = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36 QBCore/4.0.1316.400 QQBrowser/9.0.2524.400 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2875.116 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat`

func HttpsPost(mpUrl string, session *string, content *string) string {
	var ret string
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", mpUrl, strings.NewReader(*content))
	if req != nil {
		if session != nil {
			req.Header.Set("Cookie", *session)
		}
		req.Header.Set("User-Agent", UA)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, _ := client.Do(req)
		if resp != nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			ret = string(bytes)
		}
	}
	return ret
}
func HttpsGet(mpUrl string, session *string) string {
	var ret string
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", mpUrl, nil)
	if req != nil {
		if session != nil {
			req.Header.Set("Cookie", *session)
		}
		req.Header.Set("User-Agent", UA)
		resp, _ := client.Do(req)
		if resp != nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			ret = string(bytes)
		}
	}
	return ret
}

func MustA2I(str string) int {
	var ret int
	ret, _ = strconv.Atoi(str)
	return ret
}
