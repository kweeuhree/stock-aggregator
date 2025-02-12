package fetcher

import (
	"log"
)

type Fetcher struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func New(infoLog, errorLog *log.Logger) *Fetcher {
	return &Fetcher{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
