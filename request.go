package appranking

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
https://app.sensortower.com/api/android/rankings/get_category_rankings?
category=game&
country=TH&
date=2022-09-04T00:00:00.000Z&
device=MOBILE&
limit=100&
offset=0
*/

const API_URL = `https://app.sensortower.com/api/android/rankings/get_category_rankings?category=%s&country=%s&date=%s&device=MOBILE&limit=100&offset=0`

type AppData struct {
	AppId            string  `json:"app_id"`
	Country          string  `json:"canonical_country"`
	Name             string  `json:"name"`
	PublisherName    string  `json:"publisher_name"`
	Icon             string  `json:"icon"`
	OS               string  `json:"os"`
	ReleaseDate      string  `json:"release_date"`
	UpdatedDate      string  `json:"updated_date"`
	Rating           float32 `json:"rating"`
	RatingCount      int32   `json:"rating_count"`
	Price            float32 `json:"price"`
	Version          string  `json:"version"`
	Downloads        int32   `json:"downloads"`
	Revenue          int32   `json:"revenue"`
	SupportUrl       string  `json:"support_url"`
	WebsiteUrl       string  `json:"website_url"`
	PrivacyPolicyUrl string  `json:"privacy_policy_url"`
	PublisherEmail   string  `json:"publisher_email"`
	PublisherAddress string  `json:"publisher_address"`
	PublisherCountry string  `json:"publisher_country"`
	FeatureGraphic   string  `json:"feature_graphic"`
	ShortDescription string  `json:"short_description"`
	Rank             int32   `json:"rank"`
}

func Request(cat, country, date string) ([][]AppData, error) {
	apiUrl := fmt.Sprintf(API_URL, cat, country, date+"T00:00:00.000Z")
	fmt.Println(apiUrl)
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	var response [][]AppData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(response)
	return response, nil
}
