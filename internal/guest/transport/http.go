package transport

import (
	"net/http"

	"github.com/angryronald/guestlist/cmd/container"
	"github.com/angryronald/guestlist/internal/guest/endpoint"
	"github.com/angryronald/guestlist/internal/guest/public"
	libHttp "github.com/angryronald/guestlist/lib/net/http"
	"github.com/go-chi/chi"
	kitHttp "github.com/go-kit/kit/transport/http"
)

// MakeHttpRoute define the route for guest endpoint
func MakeHttpRoute(
	r *chi.Mux,
	opts []kitHttp.ServerOption,
) http.Handler {
	// I normally use group to separate the different url path
	// or to separate the endpoints which needs middleware or not
	r.Group(func(r chi.Router) {
		r.Post("/guest_list/{name}", addGuest(opts))
		r.Get("/guest_list", listGuest(opts))
	})

	r.Group(func(r chi.Router) {
		r.Get("/guests", listArrivedGuest(opts))
		r.Put("/guests/{name}", guestArrived(opts))
		r.Delete("/guests/{name}", guestLeaves(opts))
	})

	r.Group(func(r chi.Router) {
		r.Get("/seats_empty", countEmptySeats(opts))
	})

	return r
}

func addGuest(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		libHttp.NewHTTPServer(endpoint.AddGuest(container.Injector().Application.Guest()), libHttp.Option{
			DecodeModel: &public.AddGuestRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func listGuest(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		libHttp.NewHTTPServer(endpoint.ListGuests(container.Injector().Application.Guest()), libHttp.Option{}, opts).ServeHTTP(w, r)
	}
}

func listArrivedGuest(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		libHttp.NewHTTPServer(endpoint.ListArrivedGuests(container.Injector().Application.Guest()), libHttp.Option{}, opts).ServeHTTP(w, r)
	}
}

func guestArrived(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		libHttp.NewHTTPServer(endpoint.GuestArrived(container.Injector().Application.Guest()), libHttp.Option{
			DecodeModel: &public.GuestArrivedRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func guestLeaves(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		libHttp.NewHTTPServer(endpoint.GuestLeaves(container.Injector().Application.Guest()), libHttp.Option{
			DecodeModel: &public.GuestLeavesRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func countEmptySeats(opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		libHttp.NewHTTPServer(endpoint.CountEmptySeats(container.Injector().Application.Guest()), libHttp.Option{}, opts).ServeHTTP(w, r)
	}
}
