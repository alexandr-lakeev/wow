package emoji

import (
	"context"
	"io"

	proofOfWork "github.com/alexandr-lakeev/wow/internal/pkg/proof_of_work"
	"github.com/alexandr-lakeev/wow/internal/pkg/proof_of_work/dto"
	"github.com/alexandr-lakeev/wow/internal/pkg/protocol"
	"github.com/alexandr-lakeev/wow/internal/pkg/protocol/emoji/message"
	"github.com/alexandr-lakeev/wow/internal/pkg/quotes"
)

type server struct {
	hashcash proofOfWork.Verifier
	quotes   *quotes.Quotes
	logger   protocol.Logger
}

func NewServer(hashcash proofOfWork.Verifier, quotes *quotes.Quotes, logger protocol.Logger) *server {
	return &server{
		hashcash: hashcash,
		quotes:   quotes,
		logger:   logger,
	}
}

// Handle handles messages from clients
func (s *server) Handle(ctx context.Context, transport io.ReadWriter) error {
	var challenge *dto.Challenge

	for {
		msg, err := message.ReceiveMsg(transport)
		if err != nil {
			return err
		}

		s.logger.Info("I received: " + msg)

		switch {
		case message.IsClientHelloMsg(msg):
			if challenge, err = s.handleHello(transport); err != nil {
				return err
			}
		case message.IsClientSolveMsg(msg):
			if err = s.handleSolve(transport, challenge, msg); err != nil {
				return err
			}
		default:
			if err = message.SendServerWrongMsg(transport); err != nil {
				return err
			}
		}
	}
}

func (s *server) handleHello(transport io.ReadWriter) (*dto.Challenge, error) {
	challenge := s.hashcash.GetChallenge("todo")

	if err := message.SendServerHelloMsg(transport, challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}

func (s *server) handleSolve(transport io.ReadWriter, challenge *dto.Challenge, msg string) error {
	if challenge == nil {
		return message.SendServerNoChallengeMsg(transport)
	}

	counter, err := message.GetSolutionFromMsg(msg)
	if err != nil {
		return err
	}

	challenge.SetCounter(counter)

	if s.hashcash.Verify(challenge) {
		s.logger.Info("I checked solution")

		if err := message.SendServerQuoteMsg(transport, s.quotes.Get()); err != nil {
			return err
		}

		return nil
	}

	if err := message.SendServerWrongSolutionMsg(transport); err != nil {
		return err
	}

	return protocol.ErrWrongSolution
}
