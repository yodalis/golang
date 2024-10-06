package webapp

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/yodalis/labs/googlecloud/internal/temperature"
)

type Res struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP is necessary", http.StatusBadRequest)
		return
	}

	cepRegex := `^\d{8}$`
	matched, err := regexp.MatchString(cepRegex, cep)
	if err != nil {
		http.Error(w, "error validating param", http.StatusInternalServerError)
		return
	}
	if !matched {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := temperature.GetCity(cep)
	if err != nil {
		if err.Error() == temperature.ErrNotFoundZipcode.Error() {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempC := temperature.GetWeatherTemperature(city)
	tempF := temperature.ConvertToFahrenheit(tempC)
	tempK := temperature.ConvertToKelvin(tempC)

	res := Res{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
