package domain

import "time"

type DTOCreateBlog struct {
	Title       string     `json:"title" binding:"required"`
	Slug        string     `json:"slug" binding:"required"`
	Content     string     `json:"content" binding:"required"`
	Summary     *string    `json:"summary,omitempty"`
	Status      string     `json:"status" binding:"required"`
	AuthorID    *string    `json:"author_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type DTOBlogResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	Summary     *string    `json:"summary,omitempty"`
	Status      string     `json:"status"`
	AuthorID    *string    `json:"author_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func NewDTOBlogResponse(blog *Blog) *DTOBlogResponse {
	return &DTOBlogResponse{
		ID:          blog.ID.String(),
		Title:       blog.Title,
		Slug:        blog.Slug,
		Content:     blog.Content,
		Summary:     blog.Summary,
		Status:      blog.Status,
		AuthorID:    blog.AuthorID,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
}
