package webapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type RequestData struct {
	CEP string
}

type WeatherResponse struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

func CEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	tr := otel.Tracer("service-cep")

	ctx, span := tr.Start(r.Context(), "Receive CEP")
	defer span.End()

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	cepRegex := `^\d{8}$`
	matched, err := regexp.MatchString(cepRegex, data.CEP)
	if err != nil {
		http.Error(w, "error validating param", http.StatusInternalServerError)
		return
	}
	if !matched {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://goapp2:9090/weather?cep=%s", data.CEP), nil)

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao chamar o Serviço B:", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		http.Error(w, string(body), res.StatusCode)
		return
	}

	var weatherResponse WeatherResponse
	if err := json.NewDecoder(res.Body).Decode(&weatherResponse); err != nil {
		http.Error(w, "Erro ao decodificar a resposta do Serviço B", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherResponse)
}
