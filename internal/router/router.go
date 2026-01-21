package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: Initialize handlers and register routes
	return mux
}
