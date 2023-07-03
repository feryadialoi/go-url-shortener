package repository

import (
	"context"
	"database/sql"

	"github.com/feryadialoi/go-url-shortener/internal/model"
	"github.com/feryadialoi/go-url-shortener/internal/port"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type URLShortenRepository struct {
	db *sql.DB
	nr *newrelic.Application
}

func NewURLShortenRepository(db *sql.DB, nr *newrelic.Application) port.URLShortenRepository {
	return &URLShortenRepository{
		db: db,
		nr: nr,
	}
}

func (repo *URLShortenRepository) Save(ctx context.Context, us *model.URLShortenEntity) (*model.URLShortenEntity, error) {
	txn := newrelic.FromContext(ctx)

	segment := txn.StartSegment("URLShortenRepository.Save")
	defer segment.End()

	q := "INSERT INTO url_shorten (short_path, real_url) VALUES ($1, $2) RETURNING id"

	r, err := repo.db.QueryContext(ctx, q, us.ShortPath, us.RealURL)
	if err != nil {
		txn.NoticeError(err)
		return nil, err
	}

	if !r.Next() {
		txn.NoticeError(sql.ErrNoRows)
		return nil, sql.ErrNoRows
	}

	if err = r.Scan(&us.ID); err != nil {
		txn.NoticeError(err)
		return nil, err
	}

	return us, nil
}

func (repo *URLShortenRepository) FindByShortPath(ctx context.Context, shortPath string) (*model.URLShortenEntity, error) {
	txn := newrelic.FromContext(ctx)

	segment := txn.StartSegment("URLShortenRepository.FindByShortPath")
	defer segment.End()

	q := "SELECT * FROM url_shorten WHERE short_path = $1 LIMIT 1"

	r, err := repo.db.QueryContext(ctx, q, shortPath)
	if err != nil {
		txn.NoticeError(err)
		return nil, err
	}

	if !r.Next() {
		txn.NoticeError(sql.ErrNoRows)
		return nil, sql.ErrNoRows
	}

	var us model.URLShortenEntity
	err = r.Scan(&us.ID, &us.ShortPath, &us.RealURL)
	if err != nil {
		txn.NoticeError(err)
		return nil, err
	}

	return &us, nil
}
