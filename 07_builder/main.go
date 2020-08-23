package main

import (
	"./builder"
)

func main() {
	var saltWaterBuilder builder.SaltWaterInterface
	saltWaterBuilder = builder.CreateSaltWaterBuilder()

	director := builder.CreateDirector(saltWaterBuilder)
	director.Construct()

	director.Show()
}
