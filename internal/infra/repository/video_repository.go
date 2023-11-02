package repository

import (
	"api-catalogo/internal/application/dto"
	"api-catalogo/internal/domain/entity"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type VideoRepository struct {
	ElaticClient *elasticsearch.Client
}

func NewVideoRepository(client *elasticsearch.Client) *VideoRepository {
	return &VideoRepository{
		ElaticClient: client,
	}
}

func (repo *VideoRepository) List(ctx context.Context,input dto.ListVideoInputDto) ([]entity.Video,error) {
	//var mapResp map[string]any
	var searchResp dto.SearchResponse
	query := `{"query": {"match_all" : {}}}`
	sort := fmt.Sprintf(`"%v": "order": "asc"`,input.Sort)
	res, err := repo.ElaticClient.Search(
		repo.ElaticClient.Search.WithContext(ctx),
		repo.ElaticClient.Search.WithIndex("videos"),
		repo.ElaticClient.Search.WithBody(strings.NewReader(query)),
		repo.ElaticClient.Search.WithTrackTotalHits(true),
		//repo.ElaticClient.Search.WithPretty(),
		repo.ElaticClient.Search.WithSize(input.PerPage),
		repo.ElaticClient.Search.WithFrom(0),
		repo.ElaticClient.Search.WithSort(sort),
	)
	defer res.Body.Close()
	if err != nil {
		return nil,err
	}

	//if err := json.NewDecoder(res.Body).Decode(&mapResp); err != nil {
	//	log.Fatalf("Error parsing the response body: %s", err)
	//	return nil,err
	//}
	//fmt.Println(mapResp)
	
	if err := json.NewDecoder(res.Body).Decode(&searchResp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil,err
	}

	var videos []entity.Video
	for _,video := range searchResp.Hits.Hits {
		videoEntity := entity.Video{
			ID: video.Source.ID,
			Title: video.Source.Title,
			Description: video.Source.Description,
			YearLaunched: video.Source.YearLaunched,
			Duration: video.Source.Duration,
			BannerUrl: video.Source.BannerUrl,
			VideoUrl: video.Source.VideoUrl,
			CategoriesID: video.Source.CategoriesID,
			CreatedAt: video.Source.CreatedAt,
		}
		videos = append(videos, videoEntity)
	}

	return videos,nil
}

func (repo *VideoRepository) GetByID(id string) (entity.Video,error) {
	var searchResp dto.SearchResponse
	query := `{"query": {"match_all" : {"id": `+id+`}}}`
	res, err := repo.ElaticClient.Search(
		repo.ElaticClient.Search.WithIndex("videos"),
		repo.ElaticClient.Search.WithBody(strings.NewReader(query)),
		repo.ElaticClient.Search.WithTrackTotalHits(true),
	)
	defer res.Body.Close()
	if err != nil {
		return entity.Video{},err
	}

	if err := json.NewDecoder(res.Body).Decode(&searchResp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return entity.Video{},err
	}

	var videoEntity entity.Video
	for _,video := range searchResp.Hits.Hits {
		videoResp := entity.Video{
			ID: video.Source.ID,
			Title: video.Source.Title,
			Description: video.Source.Description,
			YearLaunched: video.Source.YearLaunched,
			Duration: video.Source.Duration,
			BannerUrl: video.Source.BannerUrl,
			VideoUrl: video.Source.VideoUrl,
			CategoriesID: video.Source.CategoriesID,
			CreatedAt: video.Source.CreatedAt,
		}
		videoEntity = videoResp
	}

	return videoEntity,nil
}

func (repo *VideoRepository) GetByTitle(title string) (entity.Video,error){
	var searchResp dto.SearchResponse
	query := `{"query": {"match_all" : {"id": `+title+`}}}`
	res, err := repo.ElaticClient.Search(
		repo.ElaticClient.Search.WithIndex("videos"),
		repo.ElaticClient.Search.WithBody(strings.NewReader(query)),
		repo.ElaticClient.Search.WithTrackTotalHits(true),
	)
	defer res.Body.Close()
	if err != nil {
		return entity.Video{},err
	}

	if err := json.NewDecoder(res.Body).Decode(&searchResp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return entity.Video{},err
	}

	var videoEntity entity.Video
	for _,video := range searchResp.Hits.Hits {
		videoResp := entity.Video{
			ID: video.Source.ID,
			Title: video.Source.Title,
			Description: video.Source.Description,
			YearLaunched: video.Source.YearLaunched,
			Duration: video.Source.Duration,
			BannerUrl: video.Source.BannerUrl,
			VideoUrl: video.Source.VideoUrl,
			CategoriesID: video.Source.CategoriesID,
			CreatedAt: video.Source.CreatedAt,
		}
		videoEntity = videoResp
	}

	return videoEntity,nil
}