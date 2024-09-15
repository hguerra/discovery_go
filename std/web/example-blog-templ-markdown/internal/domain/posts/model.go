package posts

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content string
	Author  Author `toml:"author"`
}
