package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TronAccount struct {
	Address string `json:"address"`
	Visible bool   `json:"visible"`
}

type TronBalanceResponse struct {
	Balance float64 `json:"balance"`
}

func FetchBalance(address string) (map[string]interface{}, error) {

	url := fmt.Sprintf("https://apilist.tronscanapi.com/api/accountv2?address=%s", address)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("TRON-PRO-API-KEY", os.Getenv("TRON_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func FetchTransactions(address string, limit, start int) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://apilist.tronscanapi.com/api/transaction?sort=-timestamp&count=true&limit=%d&start=%d&address=%s", limit, start, address)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("TRON-PRO-API-KEY", os.Getenv("TRON_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
