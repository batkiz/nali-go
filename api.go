package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type LocData struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func getIpInfo(ip, lang string) {
	if lang == "zh" {
		lang = "zh-CN"
	} else if lang == "en" {
		lang = "en"
	} else {
		lang = "zh-CN"
	}

	url := fmt.Sprintf("http://ip-api.com/json/%v?lang=%v", ip, lang)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if resp.Header.Get("X-Rl") == "0" {
		println("Reached the usage limits! plz try again after %v seconds.",
			resp.Header.Get("X-Ttl"))
		os.Exit(2)
	}

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		println("error occurred")
		os.Exit(2)
	}

	var dat LocData

	json.Unmarshal(contents, &dat)
	if dat.Status != "success" {
		fmt.Println("query failed")
		os.Exit(2)
	}

	fmt.Printf("%v\t[%v @ %v, %v]", dat.Query, dat.Isp, dat.City, dat.Country)
}
