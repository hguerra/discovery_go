package server

import (
	"encoding/json"
	"log"
	"net/http"

	"example-blog/internal/domain/posts"
	"example-blog/web"
	"example-blog/web/handlers"
	"example-blog/web/pages"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("GET /assets/", fileServer)

	postsHandler := handlers.NewPostsHandler(posts.FileReader{})

	mux.Handle("GET /web", templ.Handler(pages.HelloForm("Example")))
	mux.HandleFunc("POST /hello", handlers.HelloHandler)
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
