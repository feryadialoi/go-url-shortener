package usecase

import (
	"context"

	"github.com/feryadialoi/go-url-shortener/internal/model"
	"github.com/feryadialoi/go-url-shortener/internal/port"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type URLShortenUsecase struct {
	repo port.URLShortenRepository
	nr   *newrelic.Application
}

func NewURLShortenUsecase(repo port.URLShortenRepository, nr *newrelic.Application) port.URLShortenUsecase {
	return &URLShortenUsecase{
		repo: repo,
		nr:   nr,
	}
}

func (uc *URLShortenUsecase) Create(ctx context.Context, req *model.CreateURLShortenRequest) (*model.URLShortenResponse, error) {
	txn := newrelic.FromContext(ctx)

	segment := txn.StartSegment("URLShortenUsecase.Create")
	defer segment.End()

	us := model.URLShortenEntity{
		ShortPath: req.ShortPath,
		RealURL:   req.RealURL,
	}

	_, err := uc.repo.Save(ctx, &us)
	if err != nil {
		txn.NoticeError(err)
		return nil, err
	}

	res := model.URLShortenResponse{
		ID:        us.ID,
		ShortPath: us.ShortPath,
		RealURL:   us.RealURL,
	}

	return &res, err
}

func (uc *URLShortenUsecase) GetByShortPath(ctx context.Context, shortPath string) (*model.URLShortenResponse, error) {
	txn := newrelic.FromContext(ctx)

	segment := txn.StartSegment("URLShortenUsecase.GetByShortPath")
	defer segment.End()

	us, err := uc.repo.FindByShortPath(ctx, shortPath)
	if err != nil {
		txn.NoticeError(err)
		return nil, err
	}

	res := model.URLShortenResponse{
		ID:        us.ID,
		ShortPath: us.ShortPath,
		RealURL:   us.RealURL,
	}

	return &res, err
}
