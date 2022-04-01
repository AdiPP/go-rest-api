package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type ChiRouter struct {}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &ChiRouter{}
}

func (c *ChiRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (c *ChiRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (c *ChiRouter) Serve(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}