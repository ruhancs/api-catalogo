package main

import (
	"api-catalogo/internal/application/factory"
	"api-catalogo/internal/infra/db"
	"api-catalogo/internal/infra/web"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

//checar vulnerabilidades: govulncheck ./...

func main() {
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	
	client,err := db.NewElaticConn()
	if err != nil{
		sugar.Errorw("Error to initiate elasticsearch", "error", err)
		panic(err)
	}

	listVideosUseCase := factory.ListVideosUseCaseFactory(client)
	getVideoByTitleUseCase := factory.GetVideoByTitleUseCaseFactory(client)
	getVideoByIDUseCase := factory.GetVideoByIDUseCaseFactory(client)

	app := web.NewApplication(sugar,listVideosUseCase,getVideoByTitleUseCase,getVideoByIDUseCase)

	app.Server()
}