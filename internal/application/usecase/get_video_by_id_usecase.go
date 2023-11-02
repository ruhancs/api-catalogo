package usecase

import (
	"api-catalogo/internal/application/dto"
	"api-catalogo/internal/domain/gateway"
)

type GetVideoByIDUseCase struct {
	VideoRepository gateway.VideoRepositoryInterface
}

func NewGetVideoByIDUseCase(repository gateway.VideoRepositoryInterface) *GetVideoByIDUseCase {
	return &GetVideoByIDUseCase{
		VideoRepository: repository,
	}
}

func (usecase *GetVideoByIDUseCase) Execute(id string) (dto.GetVideoByIDOutputDto,error) {
	video,err := usecase.VideoRepository.GetByID(id)
	if err != nil {
		return dto.GetVideoByIDOutputDto{},err
	}

	output := dto.GetVideoByIDOutputDto{
		Title: video.Title,
		Description: video.Description,
		YearLaunched: video.YearLaunched,
		Duration: video.Duration,
		BannerUrl: video.BannerUrl,
		VideoUrl: video.VideoUrl,
		CategoriesID: video.CategoriesID,
	}

	return output,nil
}