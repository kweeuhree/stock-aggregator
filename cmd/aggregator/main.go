package main

import (
	"log"
	"os"
	"stock-aggregator/internal/aggregator"
	"stock-aggregator/internal/fetcher"

	"github.com/joho/godotenv"
)

type application struct {
	infoLog    *log.Logger
	errorLog   *log.Logger
	fetcher    *fetcher.Fetcher
	aggregator *aggregator.Aggregator
}

func main() {
	godotenv.Load()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	fetcher := fetcher.New(infoLog, errorLog)

	aggregator := aggregator.New(infoLog, errorLog)

	app := application{
		infoLog:    infoLog,
		errorLog:   errorLog,
		fetcher:    fetcher,
		aggregator: aggregator,
	}

	fetched, err := app.fetcher.Fetch()
	if err != nil {
		app.errorLog.Printf("failed scraping: %+v", err)
	}

	err = app.aggregator.Aggregate(fetched)
	if err != nil {
		app.errorLog.Printf("failed aggregating: %+v", err)
	}
}
