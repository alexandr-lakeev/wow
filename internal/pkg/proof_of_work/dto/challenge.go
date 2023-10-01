package dto

import (
	"strconv"
	"strings"
)

type Challenge struct {
	version    string
	complexity int
	time       int64
	rand       string
	resource   string
	counter    int
}

func NewChallenge(version string, complexity int, time int64, rand string, resource string, counter int) *Challenge {
	return &Challenge{
		version:    version,
		complexity: complexity,
		time:       time,
		rand:       rand,
		resource:   resource,
		counter:    counter,
	}
}

func (c *Challenge) String() string {
	return strings.Join([]string{
		c.version,
		strconv.Itoa(c.complexity),
		strconv.FormatInt(c.time, 10),
		c.rand,
		c.resource,
		strconv.Itoa(c.counter),
	}, ":")
}

func (c *Challenge) Bytes() []byte {
	return []byte(c.String())
}

func (c *Challenge) IncreaseCounter() {
	c.counter++
}

func (c *Challenge) SetCounter(counter int) {
	c.counter = counter
}

func (c *Challenge) GetCounter() int {
	return c.counter
}
