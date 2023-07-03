package port

import (
	"context"

	"github.com/feryadialoi/go-url-shortener/internal/model"
)

type URLShortenRepository interface {
	Save(ctx context.Context, shorten *model.URLShortenEntity) (*model.URLShortenEntity, error)
	FindByShortPath(ctx context.Context, shortPath string) (*model.URLShortenEntity, error)
}
