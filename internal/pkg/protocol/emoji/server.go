package emoji

import (
	"io"

	proofOfWork "github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work"
	"github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work/dto"
	"github.com/alexandr-lakeev/wow.git/internal/pkg/protocol"
	"github.com/alexandr-lakeev/wow.git/internal/pkg/protocol/emoji/message"
)

type server struct {
	hashcash proofOfWork.Verifier
}

func NewServer(hashcash proofOfWork.Hashcash) *server {
	return &server{
		hashcash: hashcash,
	}
}

// Handle handles messages from clients via transport
func (s *server) Handle(transport io.ReadWriter) error {
	var challenge *dto.Challenge

	for {
		msg, err := message.ReceiveMsg(transport)
		if err != nil {
			return err
		}

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
		if err := message.SendMsg(transport, "good"); err != nil {
			return err
		}
	}

	if err := message.SendServerWrongSolutionMsg(transport); err != nil {
		return err
	}

	return protocol.ErrWrongSolution
}
