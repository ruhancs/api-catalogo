package gateway

import (
	"api-catalogo/internal/application/dto"
	"api-catalogo/internal/domain/entity"
	"context"
)

type VideoRepositoryInterface interface {
	List(ctx context.Context,input dto.ListVideoInputDto) ([]entity.Video,error)
	GetByID(id string) (entity.Video,error)
	GetByTitle(title string) (entity.Video,error)
	//GetByCategory(categoryName string) ([]entity.Video,error)
}