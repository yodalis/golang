package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	apiUrlRequest = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	dbTimeout     = 10 * time.Millisecond
	reqTimeout    = 200 * time.Millisecond
)

type ExchangeRateResponse struct {
	USDBRL ExchangeRate `json:"USDBRL"`
}

type ExchangeRate struct {
	Bid string `json:"bid"`
}

func FetchExchangeRate() (*ExchangeRateResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, reqTimeout)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("Timeout Request!")
		return nil, ctx.Err()
	default:
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrlRequest, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var bodyJson ExchangeRateResponse
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		return nil, err
	}

	return &bodyJson, nil
}

func InsertDataDB(db *sql.DB, currencyResp *ExchangeRateResponse) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("Timeout database")
		return ctx.Err()
	default:
	}

	stmt, err := db.PrepareContext(ctx, "INSERT INTO currency_data (bid) VALUES (?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(currencyResp.USDBRL.Bid)
	if err != nil {
		return err
	}

	return nil
}
