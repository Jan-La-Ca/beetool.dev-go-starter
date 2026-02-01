package domain

import "github.com/dukk308/beetool.dev-go-starter/pkgs/base"

var (
	ErrBlogExisted = &base.DomainError{
		Message: "blog already exists",
		Code:    "INVALID_USERNAME",
	}
	ErrInvalidEmail = &base.DomainError{
		Message: "slug cannot be empty",
		Code:    "INVALID_EMAIL",
	}
	ErrInvalidRole = &base.DomainError{
		Message: "invalid blog status",
		Code:    "INVALID_ROLE",
	}
	ErrUnauthorized = &base.DomainError{
		Message: "unauthorized blog action",
		Code:    "UNAUTHORIZED",
	}
)
