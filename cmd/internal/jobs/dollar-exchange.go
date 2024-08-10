package jobs

import (
	"MathXplains/internal/domain/sqlite/repository"
	"MathXplains/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type dollarExchangeResponse struct {
	Conversions map[string]float32 `json:"conversion_rates"`
}

var (
	configRepo repository.ConfigRepository
	baseUrl    = "https://v6.exchangerate-api.com/v6/%s/latest/USD"
	client     = http.Client{}
)

func SetConfigRepo(repo *repository.ConfigRepository) {
	configRepo = *repo
}

func UpdateDollarExchange() {
	fmt.Println("Updating dollar exchange rate")
	url := fmt.Sprintf(baseUrl, os.Getenv("EXCHANGE_API_KEY"))
	data, err := makeRequest(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	brl := data.Conversions["BRL"]
	err = configRepo.Set(service.DollarExchangeRateKey, fmt.Sprintf("%.2f", brl))
	if err != nil {
		fmt.Println(err)
	}
}

func makeRequest(url string) (*dollarExchangeResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return mapExchangeResponse(resp.Body)
}

func mapExchangeResponse(body io.Reader) (*dollarExchangeResponse, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	data := dollarExchangeResponse{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
