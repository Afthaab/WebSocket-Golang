package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/webSocket.com/internal/handler"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Home))
	mux.Get("/ws", http.HandlerFunc(handler.WsEndpoint))
	return mux
}
