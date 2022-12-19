/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package http

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	httn "net/http"
	"reflect"
	"time"
)

func withDefaultTimeout(srv httn.Handler) httn.Handler {
	return httn.TimeoutHandler(srv, time.Second*30, "Request timed out.")
}

func withExtendedTimeout(interval time.Duration, srv httn.Handler) httn.Handler {
	return httn.TimeoutHandler(srv, interval, "Request timed out.")
}

func Serve(
	ep endpoint.Endpoint,
	dec http.DecodeRequestFunc,
	enc http.EncodeResponseFunc,
) httn.Handler {
	return withDefaultTimeout(http.NewServer(ep, dec, enc))
}

func ServeWithTimeout(
	ep endpoint.Endpoint,
	dec http.DecodeRequestFunc,
	enc http.EncodeResponseFunc,
	interval time.Duration,
) httn.Handler {
	return withExtendedTimeout(interval, http.NewServer(ep, dec, enc))
}

func Err(r interface{}) string {
	v := reflect.ValueOf(r)
	f := v.FieldByName("Err")
	return f.String()
}
