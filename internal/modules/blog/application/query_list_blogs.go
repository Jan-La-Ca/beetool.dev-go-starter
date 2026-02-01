package application

import (
	"context"

	"github.com/dukk308/beetool.dev-go-starter/internal/modules/blog/domain"
	"github.com/dukk308/beetool.dev-go-starter/pkgs/base"
)

type ListBlogsQuery struct {
	repository domain.IBlogRepository
}

func NewListBlogsQuery(repository domain.IBlogRepository) *ListBlogsQuery {
	return &ListBlogsQuery{
		repository: repository,
	}
}

func (q *ListBlogsQuery) Execute(ctx context.Context) ([]*domain.DTOBlogResponse, error) {
	blogs, err := q.repository.GetAll(ctx)
	if err != nil {
		return nil, base.ToDomainError(err)
	}
	result := make([]*domain.DTOBlogResponse, len(blogs))
	for i, b := range blogs {
		result[i] = domain.NewDTOBlogResponse(b)
	}
	return result, nil
}
