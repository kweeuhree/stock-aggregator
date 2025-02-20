package aggregator

type Aggregator struct {
	Highs  []interface{}
	Lows   []interface{}
	Opens  []interface{}
	Closes []interface{}
}

func New(size int) *Aggregator {
	return &Aggregator{
		Highs:  make([]interface{}, 0, size),
		Lows:   make([]interface{}, 0, size),
		Opens:  make([]interface{}, 0, size),
		Closes: make([]interface{}, 0, size),
	}
}
