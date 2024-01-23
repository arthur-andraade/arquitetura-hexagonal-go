package web

import (
	"arquitetura-hexagonal/adapters/web/handler"
	"arquitetura-hexagonal/application"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w *WebServer) Server() {

	mux := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(mux, n, w.Service)
	http.Handle("/", mux)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
