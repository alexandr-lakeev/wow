package message

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func IsMsg(msg string, check string) bool {
	return strings.IndexRune(msg, []rune(check)[0]) == 0
}

func IsClientHelloMsg(msg string) bool {
	return IsMsg(msg, MsgProtocolClientHello)
}

func IsClientSolveMsg(msg string) bool {
	return IsMsg(msg, MsgProtocolClientSolve)
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
