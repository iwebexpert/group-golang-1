package server

import (
	"github.com/go-chi/chi"
)

func (serv *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		//..
	})
}
