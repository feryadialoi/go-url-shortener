package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/feryadialoi/go-url-shortener/config"
	"github.com/feryadialoi/go-url-shortener/internal/adapter/repository"
	"github.com/feryadialoi/go-url-shortener/internal/adapter/usecase"
	"github.com/feryadialoi/go-url-shortener/internal/model"
	"github.com/feryadialoi/go-url-shortener/pkg/database"
	"github.com/feryadialoi/go-url-shortener/pkg/newrelic"
	"github.com/feryadialoi/go-url-shortener/pkg/respond"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	newrelic2 "github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := config.MustLoad()
	db := database.MustNew(conf)
	nr := newrelic.MustNew(conf)
	urlShortenRepo := repository.NewURLShortenRepository(db, nr)
	urlShortenUc := usecase.NewURLShortenUsecase(urlShortenRepo, nr)

	e := echo.New()
	e.Use(nrecho.Middleware(nr))

	e.POST("/api/v1/url-shortens", func(c echo.Context) error {
		txn := nrecho.FromContext(c)

		segment := txn.StartSegment("URLShortenHandler.CreateURLShorten")
		defer segment.End()

		ctx := newrelic2.NewContext(c.Request().Context(), txn)

		var req model.CreateURLShortenRequest
		if err := c.Bind(&req); err != nil {
			txn.NoticeError(err)
			return respond.ResponseBadRequest(c, err)
		}

		res, err := urlShortenUc.Create(ctx, &req)
		if err != nil {
			txn.NoticeError(err)
			return respond.ResponseInternalServerError(c, err)
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/:shortPath", func(c echo.Context) error {
		txn := nrecho.FromContext(c)

		segment := txn.StartSegment("URLShortenHandler.GetURLShorten")
		defer segment.End()

		ctx := newrelic2.NewContext(c.Request().Context(), txn)

		us, err := urlShortenUc.GetByShortPath(ctx, c.Param("shortPath"))
		if err != nil {
			txn.NoticeError(err)
			if errors.Is(err, sql.ErrNoRows) {
				return respond.ResponseNotFound(c, err)
			}

			return respond.ResponseInternalServerError(c, err)
		}

		return c.Redirect(http.StatusTemporaryRedirect, us.RealURL)
	})

	logrus.Fatal(e.Start(conf.Host()))
}
