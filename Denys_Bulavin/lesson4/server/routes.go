package server

import "github.com/go-chi/chi"

func (server *Server) bindRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/{template}", server.getTemplateHandler)
		r.Route("/api/v1", func(r chi.Router) {
			r.Post("/articles", server.postArticleHandler)
			r.Delete("/article/{id}", server.deleteArticleHandler)
			//			r.Put("/article/{id}", server.putArticleHandler)
		})
	})
}
