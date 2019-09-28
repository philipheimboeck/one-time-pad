package main

import (
	"os"

	"./controller"
	"./model"
	"./persistence"
)

func main() {
	redisAddress := os.Getenv("REDIS_DSN")
	repository := persistence.MakeRedisRepository(redisAddress)
	model := model.MakeDefaultModel(&repository)
	controller.Start(&model)
}
