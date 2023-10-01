package client

import (
	"context"
	"fmt"
	"net"
)

type client struct {
	host string
}

func New() *client {
	return &client{}
}

func (c *client) Run(ctx context.Context, address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	fmt.Println("connected to", address)
	defer conn.Close()

	return nil
}
