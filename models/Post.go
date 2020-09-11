package models

// Post contains id, title and content
type Post struct {
	id      string
	title   string
	content string
}

// NewPost is a contructor
func NewPost(id, title, content string) *Post {
	return &Post{
		id,
		title,
		content,
	}
}

// GetID is "id" getter
func (p Post) GetID() string {
	return p.id
}
