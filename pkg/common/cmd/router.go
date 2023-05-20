package cmd

import (
	"Apis/go/pkg/mod/github.com/go-chi/chi@v4.0.1+incompatible"
	"Apis/go/pkg/mod/github.com/go-openapi/runtime@v0.25.0/middleware"

	"github.com/go-chi/chi"
)

func CreateRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r

}
