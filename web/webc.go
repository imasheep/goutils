package web

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func SimGet(requestUrl string) (respHttpCode int, respBody string, err error) {

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	respHttpCode = resp.StatusCode
	respBody = string(body)

	return
}

func SimPost(requestUrl string, postData []byte) (respHttpCode int, respBody string, err error) {

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(string(postData)))
	if err != nil {
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	respHttpCode = resp.StatusCode
	respBody = string(body)

	return
}
