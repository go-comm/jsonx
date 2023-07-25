package benchmarks

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-comm/jsonx"
	"github.com/go-comm/jsonx/drivers/jsoniter"
	"github.com/go-comm/jsonx/drivers/sonic"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Gender    bool      `json:"gender"`
	Profile   string    `json:"profile"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

var StdUser User
var JUser User
var SUser User

var user *User
var data []byte

func init() {
	user = &User{
		ID:        1000001,
		Username:  "user",
		Nickname:  "john",
		Gender:    true,
		Profile:   "http://xx.xx.xx.xx:8000/a/b/c.jpg",
		Email:     "user@example.com",
		CreatedAt: time.Now(),
	}
	data, _ = json.Marshal(user)
}

func BenchmarkJsonUnMarshal(b *testing.B) {

	var stdjson = jsonx.StandardAPI
	b.Run("stdjson unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stdjson.Unmarshal(data, user)
		}
	})

	var jsoniter = jsoniter.JsoniterAPI
	b.Run("jsoniter unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = jsoniter.Unmarshal(data, user)
		}
	})

	var sonic = sonic.SonicAPI
	b.Run("sonic unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = sonic.Unmarshal(data, user)
		}
	})

}

func BenchmarkJsonMarshal(b *testing.B) {

	var stdjson = jsonx.StandardAPI
	b.Run("stdjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = stdjson.Marshal(user)
		}
	})

	var jsoniter = jsoniter.JsoniterAPI
	b.Run("jsoniter marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = jsoniter.Marshal(user)
		}
	})

	var sonic = sonic.SonicAPI
	b.Run("sonic marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = sonic.Marshal(user)
		}
	})

}
