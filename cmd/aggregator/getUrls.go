package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getUrls(stock string) ([]string, error) {
	godotenv.Load()

	marketStackApikey := os.Getenv("MARKET_STACK_API_KEY")
	fmpCloudApiKey := os.Getenv("FMP_CLOUD_API_KEY")
	finHubApiKey := os.Getenv("FINHUB_API_KEY")

	if marketStackApikey == "" || fmpCloudApiKey == "" {
		return nil, fmt.Errorf("empty api key")
	}

	marketStackUrl := fmt.Sprintf("https://api.marketstack.com/v1/eod/latest?access_key=%s&symbols=%s", marketStackApikey, stock)
	fmpCloudURl := fmt.Sprintf("https://fmpcloud.io/api/v3/quote/%s?apikey=%s", stock, fmpCloudApiKey)
	finHubUrl := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", stock, finHubApiKey)

	urls := []string{marketStackUrl, finHubUrl, fmpCloudURl}

	return urls, nil
}
