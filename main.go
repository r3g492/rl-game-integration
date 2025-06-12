package main

import (
	"fmt"
	"war-game-poc/input"
	"war-game-poc/output"
)

func main() {
	output.InitWindow(800, 450)
	defer output.CloseWindow()

	for !output.ShouldClose() {
		keyboardInput := input.GetKeyboardInput()
		fmt.Println(keyboardInput)
		mouseInput := input.GetMouseInput()
		fmt.Println(mouseInput)
		output.DrawOutput()
	}
}
