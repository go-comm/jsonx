package sonic

import (
	"io"

	"github.com/bytedance/sonic"
	"github.com/go-comm/jsonx"
)

var json = sonic.ConfigStd

var SonicAPI jsonx.API = &sonicLibrary{}

func init() {
	jsonx.NewDecoder = SonicAPI.NewDecoder
	jsonx.NewEncoder = SonicAPI.NewEncoder
	jsonx.Marshal = SonicAPI.Marshal
	jsonx.MarshalIndent = SonicAPI.MarshalIndent
	jsonx.MarshalToString = SonicAPI.MarshalToString
	jsonx.Unmarshal = SonicAPI.Unmarshal
	jsonx.UnmarshalFromString = SonicAPI.UnmarshalFromString
	jsonx.Valid = SonicAPI.Valid
}

func SetAPI(api sonic.API) {
	if api == nil {
		api = sonic.ConfigFastest
	}
	json = api
}

type sonicLibrary struct{}

func (*sonicLibrary) NewDecoder(r io.Reader) jsonx.Decoder {
	return json.NewDecoder(r)
}

func (*sonicLibrary) NewEncoder(w io.Writer) jsonx.Encoder {
	return json.NewEncoder(w)
}

func (*sonicLibrary) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*sonicLibrary) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (*sonicLibrary) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (*sonicLibrary) UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

func (*sonicLibrary) Valid(data []byte) bool {
	return json.Valid(data)
}

func (*sonicLibrary) MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}
