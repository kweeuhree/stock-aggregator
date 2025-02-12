package aggregator

import "log"

type Aggregator struct {
	Aggregator map[string]string
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func New(infoLog, errorLog *log.Logger) *Aggregator {
	return &Aggregator{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}

func (a *Aggregator) Aggregate([][]map[string]interface{}) error {
	a.InfoLog.Println("Attempting to aggregate...")
	return nil
}
