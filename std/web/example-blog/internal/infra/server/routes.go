package server

import (
	"encoding/json"
	"log"
	"net/http"

	"example-blog/internal/domain/posts"
	"example-blog/web"
	"example-blog/web/components"
	"example-blog/web/handlers"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/web", templ.Handler(components.HelloForm()))
	mux.HandleFunc("/hello", handlers.HelloWebHandler)

	postsHandler := handlers.NewPostsHandler(posts.FileReader{})
	mux.HandleFunc("GET /posts/{slug}", postsHandler.GetPostBySlugHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
