package main

import (
	"war-game-poc/output"
)

func main() {
	output.InitWindow(800, 450)
	defer output.CloseWindow()

	for !output.ShouldClose() {
		output.DrawOutput()
	}
}
