package http

import (
	"github.com/dukk308/beetool.dev-go-starter/internal/modules/blog/application"
	"github.com/gin-gonic/gin"
)

type Http struct {
	createBlogCommand *application.CreateBlogCommand
	getBlogQuery      *application.GetBlogQuery
	listBlogsQuery    *application.ListBlogsQuery
	updateBlogCommand *application.UpdateBlogCommand
	deleteBlogCommand *application.DeleteBlogCommand
}

func NewHttp(
	createBlogCommand *application.CreateBlogCommand,
	getBlogQuery *application.GetBlogQuery,
	listBlogsQuery *application.ListBlogsQuery,
	updateBlogCommand *application.UpdateBlogCommand,
	deleteBlogCommand *application.DeleteBlogCommand,
) *Http {
	return &Http{
		createBlogCommand: createBlogCommand,
		getBlogQuery:      getBlogQuery,
		listBlogsQuery:    listBlogsQuery,
		updateBlogCommand: updateBlogCommand,
		deleteBlogCommand: deleteBlogCommand,
	}
}

func (h *Http) RegisterRoutes(router *gin.RouterGroup) {
	blogsGroup := router.Group("/v1/blogs")
	{
		blogsGroup.POST("", h.HandlerCreateBlog())
		blogsGroup.GET("", h.HandlerListBlogs())
		blogsGroup.GET("/:id", h.HandlerGetBlogByID())
		blogsGroup.GET("/slug/:slug", h.HandlerGetBlogBySlug())
		blogsGroup.PUT("/:id", h.HandlerUpdateBlog())
		blogsGroup.DELETE("/:id", h.HandlerDeleteBlog())
	}
}
