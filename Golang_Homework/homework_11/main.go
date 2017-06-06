package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bitly/go-simplejson"
)

func main() {
	c := NewClient(OpenDataUrl)
	body, err := c.GetHTTPRep()
	//resp, _ := http.Get(OpenDataUrl)
	//body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}
	var result THSRAvailSeat
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("All data is ", len(result))

	for i, v := range result {
		fmt.Println("No", i, " :", v)
	}

	js, err := simplejson.NewJson(body)
	fmt.Println(js.Get("StationID").String())
}
