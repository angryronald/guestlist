package http

import (
	"net/http"

	"github.com/angryronald/guestlist/lib/net/http/encoding"
	kitHttp "github.com/go-kit/kit/transport/http"

	guest "github.com/angryronald/guestlist/internal/guest/transport"
	"github.com/go-chi/chi"
)

func MakeHandlerV1(
	r *chi.Mux,
) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(encoding.EncodeError),
	}

	// register the endpoints from guest domain
	guest.MakeHttpRoute(r, opts)

	return r
}
