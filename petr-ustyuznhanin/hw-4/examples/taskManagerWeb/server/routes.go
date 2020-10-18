package server

import (
	"github.com/go-chi/chi"
)

func (server *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/{template}", server.getTemplateHandler)
	})
}
