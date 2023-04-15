package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

type TrueCaller struct {
	nomorCode   string
	countryCode string
}

type Phone struct {
	Carrier        string `json:"carrier"`
	NationalFormat string `json:"nationalFormat"`
}

type ResponseData struct {
	Name   string  `json:"name"`
	Phones []Phone `json:"phones"`
}

type Response struct {
	Data struct {
		Data []ResponseData `json:"data"`
	} `json:"data"`
}

func (tc *TrueCaller) fetchData() (*ResponseData, error) {
	resp, err := http.Get(fmt.Sprintf("https://xzhndvs.vercel.app/api/truecaller?nomorCode=%s&countryCode=%s", tc.nomorCode, tc.countryCode))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	return &data.Data.Data[0], nil
}

func main() {
	for {
		var nomorCode string
		var countryCode string

		prompt := &survey.Input{
			Message: "Enter nomorCode :",
		}
		survey.AskOne(prompt, &nomorCode)

		prompt = &survey.Input{
			Message: "Enter countryCode :",
		}
		survey.AskOne(prompt, &countryCode)

		tc := &TrueCaller{
			nomorCode:   nomorCode,
			countryCode: countryCode,
		}

		data, err := tc.fetchData()
		if err != nil {
			panic(err)
			os.Exit(1)
		}

		name := data.Name
		phone := data.Phones[0]
		provider := phone.Carrier
		nationalFormat := phone.NationalFormat

		result := fmt.Sprintf("Name: %s\nProvider: %s\nNational Format: %s", name, provider, nationalFormat)
		fmt.Println(result)
	}
}
