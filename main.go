package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CurrentPrice holds the api data
type CurrentPrice struct {
	BPI struct {
		USD struct {
			Code string `json:"code"`
			Rate string `json:"rate"`
		} `json:"USD"`
	} `json:"bpi"`
}

//Historical holds the api data
//type Historical struct {
//	BPI struct {
//		Previous interface{}
//	} `json:"bpi"`
//}

func main() {
	response, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		btcPrice, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		btcData := CurrentPrice{}
		err = json.Unmarshal(btcPrice, &btcData)
		fmt.Printf("The current price of bitcoin is $%s.\n", btcData.BPI.USD.Rate)
	}

	res, err := http.Get("https://api.coindesk.com/v1/bpi/historical/close.json?for=yesterday")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		histPrice, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var histData map[string]interface{}
		err = json.Unmarshal(histPrice, &histData)

		previous := histData["bpi"].(map[string]interface{})
		for _, value := range previous {
			fmt.Printf("Yesterday's price of bitcoin was $%f.\n", value.(float64))
		}
	}
}
