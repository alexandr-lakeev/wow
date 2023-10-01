package quotes

import "math/rand"

type Quote string

type Quotes struct {
	list []Quote
}

func NewService(list []Quote) *Quotes {
	return &Quotes{
		list: list,
	}
}

func (q *Quotes) Get() Quote {
	return q.list[rand.Intn(len(q.list))]
}
