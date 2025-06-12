package output

import rl "github.com/gen2brain/raylib-go/raylib"

const GameTitle = "GAME_TITLE"

func InitWindow(
	width int32,
	height int32,
) {
	rl.InitWindow(width, height, GameTitle)
	rl.SetTargetFPS(1)
}

func CloseWindow() {
	rl.CloseWindow()
}

func ShouldClose() bool {
	return rl.WindowShouldClose()
}
