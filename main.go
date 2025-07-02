package main

import (
	"fmt"
	"war-game-poc/game"
	"war-game-poc/input"
	"war-game-poc/output"
)

func main() {
	output.InitWindow(1600, 900)
	defer output.CloseWindow()

	var player = game.CreatePlayer()

	for !output.ShouldClose() {

		// gather input
		keyboardInput := input.GetKeyboardInput()
		mouseInput := input.GetMouseInput()
		fmt.Println(keyboardInput)
		fmt.Println(mouseInput)

		// update game
		game.UpdateGame(
			keyboardInput,
			mouseInput,
			player,
		)

		// draw output
		output.DrawOutput(
			player,
		)

	}
}
