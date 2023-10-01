package server

import (
	"context"
	"net"

	"github.com/alexandr-lakeev/wow/internal/pkg/protocol"
)

type server struct {
	protocol protocol.Server
}

func New(protocol protocol.Server) *server {
	return &server{
		protocol: protocol,
	}
}

func (s *server) Run(ctx context.Context, port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go s.handle(ctx, conn)
	}
}

func (s *server) handle(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	if err := s.protocol.Handle(ctx, conn); err != nil {
		return
	}
}
