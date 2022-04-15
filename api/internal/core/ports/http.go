package ports

import "net/http"

type HTTPRouter interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVEHTTP(w http.ResponseWriter, req *http.Request)
	SERVE(port string)
}
