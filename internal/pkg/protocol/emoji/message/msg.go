package message

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/alexandr-lakeev/wow/internal/pkg/protocol"
	"github.com/alexandr-lakeev/wow/internal/pkg/quotes"
)

func IsMsg(msg string, check string) bool {
	return strings.IndexRune(msg, []rune(check)[0]) == 0
}

func IsClientHelloMsg(msg string) bool {
	return IsMsg(msg, MsgProtocolClientHello)
}

func IsServerHelloMsg(msg string) bool {
	return IsMsg(msg, MsgProtocolServerHello)
}

func IsClientSolveMsg(msg string) bool {
	return IsMsg(msg, MsgProtocolClientSolve)
}

func IsQuoteMsg(msg string) bool {
	return IsMsg(msg, MsgProtocolServerQuoteMsg)
}

func GetChallengeFromMsg(msg string) (string, error) {
	if !IsServerHelloMsg(msg) {
		return "", protocol.ErrGetWrongMessage
	}

	challenge, _ := strings.CutPrefix(msg, MsgProtocolServerHello)

	return strings.TrimSpace(challenge), nil
}

func GetQuoteFromMsg(msg string) (quotes.Quote, error) {
	if !IsQuoteMsg(msg) {
		return "", protocol.ErrGetWrongMessage
	}

	quote, _ := strings.CutPrefix(msg, MsgProtocolServerQuoteMsg)

	return quotes.Quote(strings.TrimSpace(quote)), nil
}

func GetSolutionFromMsg(msg string) (int, error) {
	solution, _ := strings.CutPrefix(msg, MsgProtocolClientSolve)

	counterBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(solution))
	if err != nil {
		return 0, err
	}

	counter, err := strconv.Atoi(string(counterBytes))
	if err != nil {
		return 0, err
	}

	return counter, nil
}

func SendClientHelloMsg(w io.Writer) error {
	return SendMsg(w, MsgProtocolClientHello)
}

func SendClientSolveMsg(w io.Writer, solution int) error {
	base64Solution := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(solution)))

	return SendMsg(w, MsgProtocolClientSolve, base64Solution)
}

func SendServerHelloMsg(w io.Writer, challenge fmt.Stringer) error {
	return SendMsg(w, MsgProtocolServerHello, challenge.String())
}

func SendServerWrongMsg(w io.Writer) error {
	return SendMsg(w, MsgProtocolServerWrongMsg)
}

func SendServerNoChallengeMsg(w io.Writer) error {
	return SendMsg(w, MsgProtocolServerNoChallenge)
}

func SendServerWrongSolutionMsg(w io.Writer) error {
	return SendMsg(w, MsgProtocolServerWrongSolution)
}

func SendServerQuoteMsg(w io.Writer, quote quotes.Quote) error {
	return SendMsg(w, MsgProtocolServerQuoteMsg, string(quote))
}
