package fetcher

import (
	"log"
)

type RequesterParser interface {
	RequestAndParse(url string) ([]map[string]interface{}, error)
}

type Fetcher struct {
	ErrorLog        *log.Logger
	Urls            []string
	RequesterParser RequesterParser
}

func New(requesterParser RequesterParser, errorLog *log.Logger, urls []string) *Fetcher {
	return &Fetcher{
		ErrorLog:        errorLog,
		Urls:            urls,
		RequesterParser: requesterParser,
	}
}
