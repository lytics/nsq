package protocol

import (
	"io"
)

type StatefulReadWriter interface {
	io.ReadWriter
	GetState() int
	SetState(state int)
	String() string
}

type Protocol interface {
	IOLoop(client StatefulReadWriter) error
	Execute(client StatefulReadWriter, params ...string) ([]byte, error)
}

var Protocols = map[int32]Protocol{}

type ClientError struct {
	errStr string
}

func (e ClientError) Error() string {
	return e.errStr
}

var ClientErrBadProtocol = ClientError{"E_BAD_PROTOCOL"}