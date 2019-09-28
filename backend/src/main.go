package main

import (
	"./controller"
	"./model"
	"./persistence"
)

func main() {
	repository := persistence.MakeRedisRepository("localhost:6379")
	model := model.MakeDefaultModel(&repository)
	controller.Start(&model)
}
