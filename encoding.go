package jsonx

import (
	"io"
)

var (
	JSONEncoding Encoding = &encoding{}
)

type Encoding interface {
	Encode(w io.Writer, v interface{}) error
	Decode(r io.Reader, v interface{}) error
}

type encoding struct{}

func (e *encoding) Encode(w io.Writer, v interface{}) error {
	return NewEncoder(w).Encode(v)
}

func (e *encoding) Decode(r io.Reader, v interface{}) error {
	return NewDecoder(r).Decode(v)
}
