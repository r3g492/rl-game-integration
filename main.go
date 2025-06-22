package main

import (
	"fmt"
	"war-game-poc/game"
	"war-game-poc/input"
	"war-game-poc/output"
)

func main() {
	output.InitWindow(800, 450)
	defer output.CloseWindow()

	for !output.ShouldClose() {

		// gather input
		keyboardInput := input.GetKeyboardInput()
		mouseInput := input.GetMouseInput()
		fmt.Println(keyboardInput)
		fmt.Println(mouseInput)

		// update game
		game.UpdateGame(keyboardInput, mouseInput)

		// draw output
		output.DrawOutput()

	}
}
