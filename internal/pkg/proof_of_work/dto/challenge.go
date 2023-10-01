package dto

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

func NewChallengeFromString(challenge string) (*Challenge, error) {
	parts := strings.Split(challenge, ":")

	if len(parts) < 6 {
		return nil, errors.New("wrong challenge string")
	}

	complexity, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.Wrap(err, "wrong complexity")
	}

	time, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "wrong time")
	}

	counterBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(parts[5]))
	if err != nil {
		return nil, errors.Wrap(err, "wrong counter is not base64")
	}

	counter, err := strconv.Atoi(string(counterBytes))
	if err != nil {
		return nil, errors.Wrap(err, "wrong counter")
	}

	return &Challenge{
		version:    parts[0],
		complexity: complexity,
		time:       time,
		rand:       parts[3],
		resource:   parts[4],
		counter:    counter,
	}, nil
}

func (c *Challenge) String() string {
	return strings.Join([]string{
		c.version,
		strconv.Itoa(c.complexity),
		strconv.FormatInt(c.time, 10),
		c.rand,
		c.resource,
		base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(c.counter))),
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

func (c *Challenge) GetComplexity() int {
	return c.complexity
}
