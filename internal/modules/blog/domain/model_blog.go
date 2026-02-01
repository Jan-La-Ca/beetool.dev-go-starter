package domain

import (
	"time"

	common "github.com/dukk308/beetool.dev-go-starter/pkgs/base"
)

type Blog struct {
	common.BaseModel
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	Summary     *string    `json:"summary,omitempty"`
	Status      string     `json:"status"`
	AuthorID    *string    `json:"author_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

func NewBlog(title, slug, content string, summary *string, status string, authorID *string, publishedAt *time.Time) *Blog {
	return &Blog{
		BaseModel:   *common.GenerateBaseModel(),
		Title:       title,
		Slug:        slug,
		Content:     content,
		Summary:     summary,
		Status:      status,
		AuthorID:    authorID,
		PublishedAt: publishedAt,
	}
}
