package protocol

import (
	"context"
	"io"

	"github.com/alexandr-lakeev/wow/internal/pkg/quotes"
)

//go:generate mockery --with-expecter --name Client

type Client interface {
	GetQuote(ctx context.Context, conn io.ReadWriter) (quotes.Quote, error)
}
