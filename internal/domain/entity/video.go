package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID           string    `json:"id" valid:"required"`
	Title        string    `json:"title" valid:"required"`
	Description  string    `json:"description"`
	YearLaunched int       `json:"year_launched" valid:"required"`
	Duration     float64   `json:"duration"`
	BannerUrl    string    `json:"banner_url"`
	VideoUrl     string    `json:"video_url"`
	CategoriesID []string  `json:"categories_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewVideo(title string, yearLaunched int) (*Video, error) {
	video := &Video{
		Title:        title,
		YearLaunched: yearLaunched,
	}

	err := video.Validate()
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (v *Video) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}

	return nil
}
