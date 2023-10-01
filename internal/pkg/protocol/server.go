package protocol

import (
	"context"
	"io"
)

//go:generate mockery --with-expecter --name ReadWriter

type ReadWriter interface {
	io.ReadWriter
}

//go:generate mockery --with-expecter --name Server

type Server interface {
	Handle(ctx context.Context, conn io.ReadWriter) error
}
