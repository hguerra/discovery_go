package handlers

import (
	"bytes"
	"example-blog-templ-markdown/internal/domain/posts"
	"example-blog-templ-markdown/web/pages"
	"log"
	"net/http"
	"strings"

	"github.com/adrg/frontmatter"
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

	// var post posts.Post
	// post.Slug = slug

	post := posts.Post{
		Slug: slug,
	}

	rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
	if err != nil {
		http.Error(w, "error parsing frontmatter", http.StatusInternalServerError)
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
	err = mdRenderer.Convert(rest, &buf)
	if err != nil {
		http.Error(w, "error converting markdown", http.StatusInternalServerError)
		return
	}

	post.Content = buf.String()
	log.Println(post)
	component := pages.Post(post)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("error rendering in GetPostBySlugHandler: %e", err)
	}
}
