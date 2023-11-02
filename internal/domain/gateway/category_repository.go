package gateway

import "api-catalogo/internal/domain/entity"

type CategoryRepositoryInterface interface {
	FindAll() ([]entity.Category)
}