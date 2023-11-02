package usecase

import (
	"api-catalogo/internal/application/dto"
	"api-catalogo/internal/domain/gateway"
)

type GetVideoByTitleUseCase struct {
	VideoRepository gateway.VideoRepositoryInterface
}

func NewGetVideoByTitleUseCase(repository gateway.VideoRepositoryInterface) *GetVideoByTitleUseCase {
	return &GetVideoByTitleUseCase{
		VideoRepository: repository,
	}
}

func (usecase *GetVideoByTitleUseCase) Execute(title string) (dto.GetVideoByTitleOutputDto,error) {
	video,err := usecase.VideoRepository.GetByTitle(title)
	if err != nil {
		return dto.GetVideoByTitleOutputDto{},err
	}

	output := dto.GetVideoByTitleOutputDto{
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