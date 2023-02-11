package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gen2brain/beeep"
)

type Rate struct {
	Table    string `json:"table"`
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Rates    []struct {
		No            string  `json:"no"`
		EffectiveDate string  `json:"effectiveDate"`
		Mid           float64 `json:"mid"`
	} `json:"rates"`
}

func main() {
	eur()
}

func eur() {
	var t float64 = 4.9

	url := "http://api.nbp.pl/api/exchangerates/rates/a/eur/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := []byte(body)

	var currencyRate Rate
	json.Unmarshal(data, &currencyRate)

	fmt.Printf("EURO: %f \n", currencyRate.Rates[0].Mid)
	if currencyRate.Rates[0].Mid > t {
		err := beeep.Alert("EURO", fmt.Sprintf("Kurs franka większy niż próg. Cena: %f | Próg: %f", currencyRate.Rates[0].Mid, t), "assets/warning.png")
		if err != nil {
			panic(err)
		}
	}
}

func usd() {
	var t float64 = 5

	url := "http://api.nbp.pl/api/exchangerates/rates/a/usd/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := []byte(body)

	var currencyRate Rate
	json.Unmarshal(data, &currencyRate)

	fmt.Printf("USD: %f \n", currencyRate.Rates[0].Mid)

	if currencyRate.Rates[0].Mid > t {
		err := beeep.Alert("Dolar Amerykański", fmt.Sprintf("Kurs franka większy niż próg. Cena: %f | Próg: %f", currencyRate.Rates[0].Mid, t), "assets/warning.png")
		if err != nil {
			panic(err)
		}
	}
}
