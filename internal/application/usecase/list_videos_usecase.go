package usecase

import (
	"api-catalogo/internal/application/dto"
	"api-catalogo/internal/domain/gateway"
	"context"
)

type ListVideosUseCase struct {
	VideoRepository gateway.VideoRepositoryInterface
}

func NewListVideosUseCase(repository gateway.VideoRepositoryInterface) *ListVideosUseCase {
	return &ListVideosUseCase{
		VideoRepository: repository,
	}
}

func (usecase *ListVideosUseCase) Execute(ctx context.Context,input dto.ListVideoInputDto) (dto.ListVideosOutputDto, error) {
	validateInput(&input)
	videos,err := usecase.VideoRepository.List(ctx, input)
	if err != nil {
		return dto.ListVideosOutputDto{},err
	}

	output := dto.ListVideosOutputDto{
		Items:   videos,
		Page:    input.Page,
		PerPage: input.PerPage,
		//Filter:  input.Filter,
		Sort:    input.Sort,
		//SortDir: input.SortDir,
	}

	return output, nil
}

func validateInput(input *dto.ListVideoInputDto) {
	if input.Sort == "" {
		input.Sort = "created_at"
	}

	if input.Sort != "title" && input.Sort != "created_at" {
		input.Sort = "created_at"
	}

	//if input.SortDir == "" {
	//	input.Sort = "asc"
	//}

	//if input.Sort != "asc" && input.Sort != "desc" {
	//	input.Sort = "asc"
	//}

	if input.Page == 0 {
		input.Page = 1
	}

	if input.PerPage == 0 {
		input.PerPage = 10
	}
}
