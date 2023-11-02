package factory

import (
	"api-catalogo/internal/application/usecase"
	"api-catalogo/internal/infra/repository"

	"github.com/elastic/go-elasticsearch/v8"
)

func ListVideosUseCaseFactory(client *elasticsearch.Client) *usecase.ListVideosUseCase {
	repository := repository.NewVideoRepository(client)
	usecase := usecase.NewListVideosUseCase(repository)
	return usecase
}

func GetVideoByTitleUseCaseFactory(client *elasticsearch.Client) *usecase.GetVideoByTitleUseCase {
	repository := repository.NewVideoRepository(client)
	usecase := usecase.NewGetVideoByTitleUseCase(repository)
	return usecase
}

func GetVideoByIDUseCaseFactory(client *elasticsearch.Client) *usecase.GetVideoByIDUseCase {
	repository := repository.NewVideoRepository(client)
	usecase := usecase.NewGetVideoByIDUseCase(repository)
	return usecase
}

