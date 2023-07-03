package port

import (
	"context"

	"github.com/feryadialoi/go-url-shortener/internal/model"
)

type URLShortenUsecase interface {
	Create(ctx context.Context, req *model.CreateURLShortenRequest) (*model.URLShortenResponse, error)
	GetByShortPath(ctx context.Context, shortPath string) (*model.URLShortenResponse, error)
}
