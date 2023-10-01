package message

import (
	"bufio"
	"io"
	"strings"
	"time"
)

func ReceiveMsg(r io.Reader) (string, error) {
	rd := bufio.NewReader(r)

	req, err := rd.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(req), nil
}

func SendMsg(w io.Writer, msg ...string) error {
	time.Sleep(10 * time.Millisecond)

	_, err := io.WriteString(w, strings.Join(msg, " ")+"\n")
	if err != nil {
		return err
	}

	return nil
}
