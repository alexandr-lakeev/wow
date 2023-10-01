package protocol

import (
	"context"
	"io"
)

type Server interface {
	Handle(ctx context.Context, conn io.ReadWriter) error
}
