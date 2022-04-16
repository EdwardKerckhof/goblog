package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
)

var url string = "/api/v1"

type muxRouter struct {
	config    *configs.Config
	muxRouter *mux.Router
}

func NewMuXRouter(config *configs.Config) ports.HTTPRouter {
	url = fmt.Sprintf("/api/v%s", config.ApiVersion)
	return &muxRouter{
		config:    config,
		muxRouter: mux.NewRouter(),
	}
}

func (r *muxRouter) GET(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	r.muxRouter.PathPrefix(url).Subrouter().HandleFunc(uri, f).Methods(http.MethodGet)
}

func (r *muxRouter) POST(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	r.muxRouter.PathPrefix(url).Subrouter().HandleFunc(uri, f).Methods(http.MethodPost)
}

func (r *muxRouter) PUT(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	r.muxRouter.PathPrefix(url).Subrouter().HandleFunc(uri, f).Methods(http.MethodPut)
}

func (r *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	r.muxRouter.PathPrefix(url).Subrouter().HandleFunc(uri, f).Methods(http.MethodDelete)
}

func (r *muxRouter) SERVEHTTP(w http.ResponseWriter, req *http.Request) {
	r.muxRouter.ServeHTTP(w, req)
}

func (r *muxRouter) SERVE(port string) {
	// CORS
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	corsOrigins := handlers.AllowedOrigins([]string{r.config.ApiOriginsAllowed})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	r.muxRouter.Use(handlers.CORS(corsOrigins, corsHeaders, corsMethods))

	// Mux logging and panic recovery middleware
	r.muxRouter.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})
	r.muxRouter.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	fmt.Printf("Mux HTTP server running on port %s with prefix: %s\n", port, url)
	if err := http.ListenAndServe(port, r.muxRouter); err != nil {
		log.Fatal(err)
	}
}
