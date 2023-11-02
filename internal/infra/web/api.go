package web

import (
	"api-catalogo/internal/application/usecase"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Application struct {
	Sugar                  *zap.SugaredLogger
	ListVideosUseCase      *usecase.ListVideosUseCase
	GetVideoByTitleUseCase *usecase.GetVideoByTitleUseCase
	GetVideoByIDUseCase    *usecase.GetVideoByIDUseCase
}

func NewApplication(
	logger *zap.SugaredLogger,
	listVideosUseCase *usecase.ListVideosUseCase,
	getVideoByTitleUseCase *usecase.GetVideoByTitleUseCase,
	getVideoByIDUseCase *usecase.GetVideoByIDUseCase,
) *Application {
	return &Application{
		Sugar:                  logger,
		ListVideosUseCase:      listVideosUseCase,
		GetVideoByTitleUseCase: getVideoByTitleUseCase,
		GetVideoByIDUseCase:    getVideoByIDUseCase,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8001",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.Sugar.Infow("Runing server", "port", 8001)
	return srv.ListenAndServe()
}
