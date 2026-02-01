package application

import (
	"context"

	"github.com/dukk308/beetool.dev-go-starter/internal/modules/blog/domain"
	"github.com/dukk308/beetool.dev-go-starter/pkgs/base"
)

type CreateBlogCommand struct {
	repository domain.IBlogRepository
}

func NewCreateBlogCommand(repository domain.IBlogRepository) *CreateBlogCommand {
	return &CreateBlogCommand{
		repository: repository,
	}
}

func (c *CreateBlogCommand) Execute(ctx context.Context, dto *domain.DTOCreateBlog) (*domain.DTOBlogResponse, error) {
	blog := domain.NewBlog(dto.Title, dto.Slug, dto.Content, dto.Summary, dto.Status, dto.AuthorID, dto.PublishedAt)

	if _, err := c.repository.GetBySlug(ctx, dto.Slug); err != nil {
		return nil, err
	}

	if err := c.repository.Create(ctx, blog); err != nil {
		return nil, base.ToDomainError(err)
	}
	return domain.NewDTOBlogResponse(blog), nil
}
