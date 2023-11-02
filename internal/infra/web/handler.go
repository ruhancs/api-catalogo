package web

import (
	"api-catalogo/internal/application/dto"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func(app *Application) ListVideosHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.ListVideoInputDto
	queryPage := r.URL.Query().Get("page")
	page,err := strconv.Atoi(queryPage)
	if err != nil {
		page = 1
	}
	sort := r.URL.Query().Get("sort")
	perPage := r.URL.Query().Get("limit")
	limit,err := strconv.Atoi(perPage)
	if err != nil {
		limit = 10
	}
	inputDto.Page = page
	inputDto.PerPage = limit
	inputDto.Sort = sort

	output,err := app.ListVideosUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.Sugar.Errorw("Error to execute ListVideoshandler", "error", err)
		app.errorJson(w,err, http.StatusInternalServerError)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetVideoByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	
	output,err := app.GetVideoByIDUseCase.Execute(id)
	if err != nil {
		app.Sugar.Errorw("Error to execute GetVideoByIDHandler", "error", err)
		app.errorJson(w,err, http.StatusInternalServerError)
		return
	}
	
	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetVideoByTitleHandler(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r,"title")
	
	output,err := app.GetVideoByTitleUseCase.Execute(title)
	if err != nil {
		app.Sugar.Errorw("Error to execute GetVideoByTitleHandler", "error", err)
		app.errorJson(w,err, http.StatusInternalServerError)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}