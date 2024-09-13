package handlers

import (
	"example-blog/internal/domain/posts"
	"fmt"
	"net/http"
)

type PostsHandler struct {
	sl posts.SlugReader
}

func NewPostsHandler(sl posts.SlugReader) *PostsHandler {
	return &PostsHandler{sl: sl}
}

func (p *PostsHandler) GetPostBySlugHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	postMarkdown, err := p.sl.Read(slug)
	if err != nil {
		http.Error(w, "post no found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, postMarkdown)
}
