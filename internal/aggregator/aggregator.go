package aggregator

import "log"

type Aggregator struct {
	Aggregator *map[string]string
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func New(infoLog, errorLog *log.Logger) *Aggregator {
	return &Aggregator{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
