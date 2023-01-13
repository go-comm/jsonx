package jsonx

import (
	"encoding/json"
	"io"
)

var _ API = (*standardLibrary)(nil)

type standardLibrary struct{}

func (*standardLibrary) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}

func (*standardLibrary) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (*standardLibrary) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*standardLibrary) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (*standardLibrary) MarshalToString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

func (*standardLibrary) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (*standardLibrary) UnmarshalFromString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

func (*standardLibrary) Valid(data []byte) bool {
	return json.Valid(data)
}
