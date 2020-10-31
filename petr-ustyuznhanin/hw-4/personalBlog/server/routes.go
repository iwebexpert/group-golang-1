package server

import (
	"github.com/go-chi/chi"
)

func (serv *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/{template}", serv.getTemplateHandler)
		r.Route("api/v1", func(r chi.Router) {
			r.Post("/posts", serv.postPostHandler)
			r.Delete("/posts/{id}", serv.deletePostHandler)
			r.Put("/tasks/{id}", serv.putTaskHandler)
		})
	})
}
