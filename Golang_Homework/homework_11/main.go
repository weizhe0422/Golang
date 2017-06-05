package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("All data is ", len(result.Result.StationID))

	for _, v := range result.Result.StationID {
		fmt.Println(v)
	}
}
