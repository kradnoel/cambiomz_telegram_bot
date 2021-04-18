package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AllCurrencyResponse struct {
	Data struct {
		Currencies []struct {
			Currency string `json:"currency"`
			Buy      string `json:"buy"`
			Sell     string `json:"sell"`
		} `json:"currencies"`
	} `json:"data"`
}

type SingleCurrencyResponse struct {
	Data struct {
		Currency struct {
			Currency string `json:"currency"`
			Buy      string `json:"buy"`
			Sell     string `json:"sell"`
		} `json:"currency"`
	} `json:"data"`
}

type Currency struct {
	currency string
}

func GetCurrencies() string {
	var response AllCurrencyResponse
	var currencies string

	res, err := http.Get("https://cambiomz.herokuapp.com/api/v1.0.0/exchange?query={currencies{currency}}")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	for i, currency := range response.Data.Currencies {
		currencies = currencies + currency.Currency

		if (i + 1) < len(response.Data.Currencies) {
			currencies = currencies + ", "
		}
	}

	return currencies
}

//https://cambiomz.herokuapp.com/api/v1.0.0/exchange?query={currency(currency:%22usd%22){currency,buy,sell}}
