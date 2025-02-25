package main

import (
	"log"
	"os"
	"stock-aggregator/internal/aggregator"
	"stock-aggregator/internal/fetcher"
	"stock-aggregator/utils"
)

type application struct {
	errorLog   *log.Logger
	fetcher    *fetcher.Fetcher
	aggregator *aggregator.Aggregator
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	stock := getUserInput()

	urls, err := getUrls(stock)
	if err != nil {
		errorLog.Fatalf("failed getting urls: %s", err)
	}

	requester := &fetcher.Requester{}
	fetcher := fetcher.New(requester, errorLog, urls)
	aggregator := aggregator.New(len(urls))

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
		return err
	}

	err = app.aggregator.Aggregate(fetched)
	if err != nil {
		app.errorLog.Printf("failed aggregating: %+v", err)
	}

	result := app.aggregator.CalculateAverages()

	infoLog.Println("Aggregated averages:")
	utils.PrettyPrint(result)

	return nil
}
