package server

import "github.com/go-chi/chi"

func (server *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/{template}", server.getTemplateHandler)
		r.Route("/api/v1", func(r chi.Router) {
			r.Post("/posts", server.NewPostHandler)
			r.Delete("/posts/{id}", server.deletePostHandler)
			r.Put("/posts/{id}", server.putPostHandler)
		})
	})
}
