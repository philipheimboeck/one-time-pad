package main

import (
	"./controller"
	"./model"
)

func main() {
	model := model.DefaultModel{}
	controller.Start(model)
}
