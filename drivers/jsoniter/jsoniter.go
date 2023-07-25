package jsoniter

import (
	"io"

	"github.com/go-comm/jsonx"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var JsoniterAPI jsonx.API = &jsoniterLibrary{}

func init() {
	jsonx.NewDecoder = JsoniterAPI.NewDecoder
	jsonx.NewEncoder = JsoniterAPI.NewEncoder
	jsonx.Marshal = JsoniterAPI.Marshal
	jsonx.MarshalIndent = JsoniterAPI.MarshalIndent
	jsonx.MarshalToString = JsoniterAPI.MarshalToString
	jsonx.Unmarshal = JsoniterAPI.Unmarshal
	jsonx.UnmarshalFromString = JsoniterAPI.UnmarshalFromString
	jsonx.Valid = JsoniterAPI.Valid
}

type jsoniterLibrary struct{}

func (*jsoniterLibrary) NewDecoder(r io.Reader) jsonx.Decoder {
	return json.NewDecoder(r)
}

func (*jsoniterLibrary) NewEncoder(w io.Writer) jsonx.Encoder {
	return json.NewEncoder(w)
}

func (*jsoniterLibrary) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*jsoniterLibrary) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (*jsoniterLibrary) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (*jsoniterLibrary) UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

func (*jsoniterLibrary) Valid(data []byte) bool {
	return json.Valid(data)
}

func (*jsoniterLibrary) MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}
