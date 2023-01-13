package jsonx

import (
	"encoding/json"
	"io"
)

var _ = json.Marshal

type Decoder interface {
	Buffered() io.Reader
	Decode(v interface{}) error
	DisallowUnknownFields()
	More() bool
	UseNumber()
}

type Encoder interface {
	Encode(v interface{}) error
	SetEscapeHTML(on bool)
	SetIndent(prefix, indent string)
}

type API interface {
	NewDecoder(r io.Reader) Decoder
	NewEncoder(w io.Writer) Encoder
	Marshal(v interface{}) ([]byte, error)
	MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	MarshalToString(v interface{}) (string, error)
	Unmarshal(data []byte, v interface{}) error
	UnmarshalFromString(str string, v interface{}) error
	Valid(data []byte) bool
}

var Standard API = &standardLibrary{}

var NewDecoder func(r io.Reader) Decoder = Standard.NewDecoder

var NewEncoder func(w io.Writer) Encoder = Standard.NewEncoder

var Marshal func(v interface{}) ([]byte, error) = Standard.Marshal

var MarshalIndent func(v interface{}, prefix, indent string) ([]byte, error) = Standard.MarshalIndent

var MarshalToString func(v interface{}) (string, error) = Standard.MarshalToString

var Unmarshal func(data []byte, v interface{}) error = Standard.Unmarshal

var UnmarshalFromString func(str string, v interface{}) error = Standard.UnmarshalFromString

var Valid func(data []byte) bool = Standard.Valid
