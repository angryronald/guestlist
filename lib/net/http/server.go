package http

import (
	"github.com/angryronald/guestlist/lib/net/http/decoding"
	"github.com/angryronald/guestlist/lib/net/http/encoding"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	gokitHttp "github.com/go-kit/kit/transport/http"
)

type Option struct {
	DecodeModel interface{}
	Encoder     gokitHttp.EncodeResponseFunc
	Decoder     gokitHttp.DecodeRequestFunc
}

func NewHTTPServer(
	endpoint endpoint.Endpoint,
	option Option,
	serverOption []http.ServerOption) *http.Server {
	if option.Encoder == nil {
		option.Encoder = encoding.Encode()
	}
	if option.Decoder == nil {
		option.Decoder = decoding.Decode(option.DecodeModel)
	}

	return http.NewServer(endpoint, option.Decoder, option.Encoder, serverOption...)
}
