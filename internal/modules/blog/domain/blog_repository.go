package domain

import "context"

type IBlogRepository interface {
	GetByID(ctx context.Context, id string) (*Blog, error)
	GetBySlug(ctx context.Context, slug string) (*Blog, error)
	GetAll(ctx context.Context) ([]*Blog, error)
	Create(ctx context.Context, blog *Blog) error
	Update(ctx context.Context, blog *Blog) error
	Delete(ctx context.Context, id string) error
}
