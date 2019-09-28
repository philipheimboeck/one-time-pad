package main

import (
	"os"

	"./controller"
	"./model"
	"./persistence"
)

func main() {
	redisAddress := os.Getenv("REDIS_DSN")
	port := os.Getenv("HTTP_PORT")
	repository := persistence.MakeRedisRepository(redisAddress)
	model := model.MakeDefaultModel(&repository)
	controller.Start(port, &model)
}
