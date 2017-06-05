package main

import (
	"io/ioutil"

	"net/http"
)

type client struct {
	url string
}

func NewClient(urlName string) *client {
	c := new(client)
	c.url = urlName
	//log.Fatal(c.url)
	return c
}

func (c *client) GetHTTPRep() ([]byte, error) {
	return GetHTTPReponse(c.url)

}

func GetHTTPReponse(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	reponse, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(reponse.Body)
	defer reponse.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil

}
