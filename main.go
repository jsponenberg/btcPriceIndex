package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//JSONData holds the api data
type JSONData struct {
	BPI struct {
		USD struct {
			Code string `json:"code"`
			Rate string `json:"rate"`
		} `json:"USD"`
	} `json:"bpi"`
}

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

		btcData := JSONData{}
		err = json.Unmarshal(btcPrice, &btcData)
		fmt.Printf("The current price of bitcoin is $%s.\n", btcData.BPI.USD.Rate)
	}
}
