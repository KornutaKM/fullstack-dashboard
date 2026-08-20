package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type NewsService struct {
	apiKey string
	client *http.Client
}

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
	} `json:"articles"`
}

func NewNewsService(apiKey string) *NewsService {
	return &NewsService{
		apiKey: apiKey,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (s *NewsService) GetTopHeadlines(ctx context.Context, country, category string) (*NewsResponse, error) {
	//Формируем URL
	baseUrl := "https://newsapi.org/v2/top-headlines"
	params := url.Values{}
	params.Add("apikey", s.apiKey)
	params.Add("country", country)
	params.Add("category", category)
	params.Add("pageSize", "5")

	fullURL := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	//Создаем запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	//Выполняем запрос
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create fetch news: %w", err)
	}
	defer resp.Body.Close()

	//Проверяем статус
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("news API returnes status: %d", resp.StatusCode)
	}

	//Декодируем JSON
	var result NewsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode news response: %w", err)
	}

	return &result, nil
}

func (s *NewsService) GetEverything(ctx context.Context, query string, language string) (*NewsResponse, error) {
	baseURL := "https://newsapi.org/v2/everything"
	params := url.Values{}
	params.Add("apiKey", s.apiKey)
	params.Add("q", query)
	params.Add("language", language)
	params.Add("pageSize", "5")
	params.Add("sortBy", "publishedAt")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("news API returned status: %d", resp.StatusCode)
	}

	var result NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode news response: %w", err)
	}

	return &result, nil
}
