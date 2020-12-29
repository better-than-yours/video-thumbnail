// Package api handle work with api
package api

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

// Server provides HTTP API
type Server struct {
	Store      *gorm.DB
	httpServer *http.Server
}

// Run starts http server for API with all routes
func (s *Server) Run(port int) {
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "OK")
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		img, err := GetImage(url)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		err = png.Encode(w, img)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})
	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("[ERROR] start http server, %s", err)
	}
}
