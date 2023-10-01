package protocol

import "io"

type Server interface {
	Handle(conn io.ReadWriter) error
}
