package temperature

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
)

const (
	API_KEY          = "fd1b74e1200c4bceb83215446241409"
	BASE_URL         = "http://api.weatherapi.com/v1/current.json"
	BASE_URL_VIA_CEP = "http://viacep.com.br/ws/"
)

var (
	ErrNotFoundZipcode = errors.New("can not find zipcode")
)

type ViaCEP struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetCity(cep string) (string, error) {
	res, err := http.Get(BASE_URL_VIA_CEP + cep + "/json/")
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		if res.StatusCode == http.StatusBadRequest {
			return "", ErrNotFoundZipcode
		}
		return "", err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return "", err
	}
	if c.Erro != "" {
		return "", ErrNotFoundZipcode
	}

	return c.Localidade, nil
}

func GetWeatherTemperature(city string) float64 {
	city = url.QueryEscape(city)
	url := fmt.Sprintf("%s?key=%s&q=%s", BASE_URL, API_KEY, city)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error to do request:", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading responde:", err)
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Println("error parsing res:", err)
	}

	return weather.Current.TempC
}

func ConvertToFahrenheit(celsiusTemp float64) float64 {
	fahrenheitTemp := celsiusTemp*1.8 + 32

	finalTemp := math.Round(fahrenheitTemp*10) / 10

	return finalTemp
}

func ConvertToKelvin(celsiusTemp float64) float64 {
	kelvinTemp := celsiusTemp + 273

	finalTemp := math.Round(kelvinTemp*10) / 10

	return finalTemp
}
