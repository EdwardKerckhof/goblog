package ports

import "net/http"

type PostHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}
