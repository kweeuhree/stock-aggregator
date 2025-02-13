package main

import (
	"log"
	"os"
	"stock-aggregator/internal/aggregator"
	"stock-aggregator/internal/fetcher"
)

type application struct {
	errorLog   *log.Logger
	fetcher    *fetcher.Fetcher
	aggregator *aggregator.Aggregator
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	urls, err := getUrls()
	if err != nil {
		errorLog.Fatalf("failed getting urls: %s", err)
	}

	requester := &fetcher.Requester{}
	fetcher := fetcher.New(requester, errorLog, urls)
	aggregator := aggregator.New(errorLog)

	app := application{
		errorLog:   errorLog,
		fetcher:    fetcher,
		aggregator: aggregator,
	}

	if err := app.run(); err != nil {
		errorLog.Fatal(err)
	}
}

func (app *application) run() error {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	infoLog.Println("Fetching data concurrently...")
	fetched, err := app.fetcher.Fetch()
	if err != nil {
		app.errorLog.Printf("failed fetching: %+v", err)
	}

	err = app.aggregator.Aggregate(fetched)
	if err != nil {
		app.errorLog.Printf("failed aggregating: %+v", err)
	}

	return nil
}
