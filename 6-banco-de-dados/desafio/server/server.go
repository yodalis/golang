package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

func main() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

func fetchExchangeRate() (*ExchangeRateResponse, error) {
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

func createTableDB(db *sql.DB) error {
	sql := `CREATE TABLE currency_data (
		id SERIAL PRIMARY KEY,
		bid NUMERIC
	);
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func insertDataDB(db *sql.DB, currencyResp *ExchangeRateResponse) error {
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

func handler(w http.ResponseWriter, r *http.Request) {
	currencyResp, err := fetchExchangeRate()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "./currency_data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = createTableDB(db)
	if err != nil {
		log.Fatal(err)
	}

	err = insertDataDB(db, currencyResp)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(currencyResp.USDBRL.Bid)
}
