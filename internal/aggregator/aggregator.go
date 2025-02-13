package aggregator

import "log"

type Aggregator struct {
	Aggregator *map[string]string
	ErrorLog   *log.Logger
}

func New(errorLog *log.Logger) *Aggregator {
	return &Aggregator{
		ErrorLog: errorLog,
	}
}
