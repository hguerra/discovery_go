package handlers

import (
	"bytes"
	"example-blog/internal/domain/posts"
	"example-blog/web/pages"
	"log"
	"net/http"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
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

	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)

	var buf bytes.Buffer
	err = mdRenderer.Convert([]byte(postMarkdown), &buf)
	if err != nil {
		http.Error(w, "error converting markdown", http.StatusInternalServerError)
		return
	}

	component := pages.Post(
		"My First Post",
		"Heitor Carneiro",
		buf.String(),
	)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("error rendering in GetPostBySlugHandler: %e", err)
	}
}
