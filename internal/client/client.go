package client

import (
	"context"
	"fmt"
	"net"

	"github.com/alexandr-lakeev/wow/internal/pkg/protocol"
)

type client struct {
	protocol protocol.Client
}

func New(protocol protocol.Client) *client {
	return &client{
		protocol: protocol,
	}
}

func (c *client) Run(ctx context.Context, address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	quote, err := c.protocol.GetQuote(ctx, conn)
	if err != nil {
		fmt.Println(err)
	}

	//
	_ = quote

	return nil
}
