package proof_of_work

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work/dto"
)

type hashcash struct {
	hasher     Hasher
	version    string
	complexity int
	prefix     string
	max        int
}

func New(hasher Hasher, version string, complexity int, prefix string, max int) *hashcash {
	return &hashcash{
		hasher:     hasher,
		version:    version,
		complexity: complexity,
		prefix:     prefix,
		max:        max,
	}
}

// GetChallenge returns a new challenge for the resource
func (h *hashcash) GetChallenge(resource string) *dto.Challenge {
	rnd := rand.Intn(100000)
	base := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(rnd)))

	return dto.NewChallenge(h.version, h.complexity, time.Now().Unix(), base, resource, 1)
}

// Solve solves the challenge with bruteforce
func (h *hashcash) Solve(challenge *dto.Challenge) (int, error) {
	for challenge.GetCounter() < h.max {
		if h.Verify(challenge) {
			return challenge.GetCounter(), nil
		}

		challenge.IncreaseCounter()
	}

	return 0, ErrNoSolutionFound
}

// Verify verifies that challenge was solved correctly
func (h *hashcash) Verify(challenge *dto.Challenge) bool {
	return h.verifyHash(h.hasher.Hash(challenge.Bytes()), challenge.GetComplexity())
}

func (h *hashcash) verifyHash(hash []byte, complexity int) bool {
	str := hex.EncodeToString(hash)
	for _, val := range str[:complexity] {
		if string(val) != h.prefix {
			return false
		}
	}

	return true
}
