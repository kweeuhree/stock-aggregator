package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getUrls() ([]string, error) {
	godotenv.Load()

	marketStackApikey := os.Getenv("MARKET_STACK_API_KEY")
	fmpCloudApiKey := os.Getenv("FMP_CLOUD_API_KEY")

	if marketStackApikey == "" || fmpCloudApiKey == "" {
		return nil, fmt.Errorf("empty api key")
	}

	stocks := "AAPL,MSFT"

	// marketStackUrl := fmt.Sprintf("https://api.marketstack.com/v1/intraday/latest?access_key=%s&symbols=%s", marketStackApikey, stocks)

	fmpCloudURl := fmt.Sprintf("https://fmpcloud.io/api/v3/quote/%s?apikey=%s", stocks, fmpCloudApiKey)

	urls := []string{fmpCloudURl}

	return urls, nil
}
