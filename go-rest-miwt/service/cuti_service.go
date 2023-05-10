package service

import (
	"bryanagamk/go-rest-miwt/web"
	"context"
)

type CutiService interface {
	Create(ctx context.Context, request web.CutiCreateRequest) web.CutiResponse
	Update(ctx context.Context, request web.CutiUpdateRequest) web.CutiResponse
	Delete(ctx context.Context, cutiId string)
	FindById(ctx context.Context, cutiId string) web.CutiResponse
	FindAll(ctx context.Context) []web.CutiResponse
}
