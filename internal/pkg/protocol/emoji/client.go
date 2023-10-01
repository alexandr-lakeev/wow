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

type client struct {
	hashcash proofOfWork.Solver
	logger   protocol.Logger
}

func NewClient(hashcash proofOfWork.Solver, logger protocol.Logger) *client {
	return &client{
		hashcash: hashcash,
		logger:   logger,
	}
}

// GetQuote gets quotes from server
func (c *client) GetQuote(ctx context.Context, transport io.ReadWriter) (quotes.Quote, error) {
	if err := message.SendClientHelloMsg(transport); err != nil {
		return "", err
	}

	msg, err := message.ReceiveMsg(transport)
	if err != nil {
		return "", err
	}

	c.logger.Info("I received: " + msg)

	challengeMsg, err := message.GetChallengeFromMsg(msg)
	if err != nil {
		return "", err
	}

	challenge, err := dto.NewChallengeFromString(challengeMsg)
	if err != nil {
		return "", err
	}

	c.logger.Info("I am solving the challenge: " + challenge.String())

	solution, err := c.hashcash.Solve(challenge)
	if err != nil {
		return "", err
	}

	c.logger.Info("I have solved a challenge")

	err = message.SendClientSolveMsg(transport, solution)
	if err != nil {
		return "", err
	}

	msg, err = message.ReceiveMsg(transport)
	if err != nil {
		return "", err
	}

	c.logger.Info("I received: " + msg)

	quote, err := message.GetQuoteFromMsg(msg)
	if err != nil {
		return "", err
	}

	c.logger.Info("My quote is: \"" + string(quote) + "\"")

	return quote, nil
}
