package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
) 

type MuxRouter struct {}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &MuxRouter{}
}

func (m *MuxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (m *MuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (m *MuxRouter) Serve(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}