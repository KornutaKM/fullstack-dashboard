package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ExchangeService struct {
	apiKey string
	client *http.Client
}

type ExchangeResponse struct {
	Result            string `json:"result"`
	Documentation     string `json:"documentation"`
	TermsOfUse        string `json:"terms_of_use"`
	TimeLastUpdate    int    `json:"time_last_update_unix"`
	TimeLastUpdateUTC string `json:"time_last_update_utc"`
	TimeNextUpdate    int    `json:"time_next_update_unix"`
	TimeNextUpdateUTC string `json:"time_next_update_utc"`
	BaseCode          string `json:"base_code"`
	ConversionRates   struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
		GBP float64 `json:"GBP"`
		RUB float64 `json:"RUB"`
	} `json:"conversion_rates"`
}

func NewExchangeService(apikey string) *ExchangeService {
	return &ExchangeService{
		apiKey: apikey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *ExchangeService) GetExchange(ctx context.Context, baseCurency string) (*ExchangeResponse, error) {
	baseUrl := "https://v6.exchangerate-api.com/v6"

	fullURL := fmt.Sprintf("%s/%s/latest/%s", baseUrl, s.apiKey, baseCurency)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("exchange API returnes status code: %d", resp.StatusCode)
	}

	var result ExchangeResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil

}
