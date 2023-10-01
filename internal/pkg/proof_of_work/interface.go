package proof_of_work

import (
	"time"

	"github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work/dto"
)

//go:generate mockery --with-expecter --name Hasher

type Hasher interface {
	Hash([]byte) []byte
}

type Hashcash interface {
	GetChallenge(resource string) *dto.Challenge
	Solve(challenge *dto.Challenge) (int, error)
	Verify(challenge *dto.Challenge) bool
}

type Solver interface {
	Solve(challenge *dto.Challenge) (int, error)
}

type Verifier interface {
	GetChallenge(resource string) *dto.Challenge
	Verify(challenge *dto.Challenge) bool
}

type Clock interface {
	Now() time.Time
}
