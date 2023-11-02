package dto

import (
	"api-catalogo/internal/domain/entity"
	"time"
)

type ListVideoInputDto struct {
	Page    int
	PerPage int
	//Filter  string
	Sort    string
	//SortDir string
}

type ListVideosOutputDto struct {
	Items   []entity.Video
	Page    int
	PerPage int
	Filter  string
	Sort    string
	//SortDir string
}

type GetVideoByIDOutputDto struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	YearLaunched int       `json:"year_launched"`
	Duration     float64   `json:"duration"`
	BannerUrl    string    `json:"banner_url"`
	VideoUrl     string    `json:"video_url"`
	CategoriesID []string  `json:"categories_id"`
}

type GetVideoByTitleOutputDto struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	YearLaunched int       `json:"year_launched"`
	Duration     float64   `json:"duration"`
	BannerUrl    string    `json:"banner_url"`
	VideoUrl     string    `json:"video_url"`
	CategoriesID []string  `json:"categories_id"`
}

type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*struct {
			Source *Video `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Video struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	YearLaunched int       `json:"year_launched"`
	Duration     float64   `json:"duration"`
	IsPublished  bool      `json:"is_published"`
	BannerUrl    string    `json:"banner_url"`
	VideoUrl     string    `json:"video_url"`
	CategoriesID []string  `json:"categories_id"`
	CreatedAt    time.Time `json:"created_at"`
}
