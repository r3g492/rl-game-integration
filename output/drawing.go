package output

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawOutput() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	rl.EndDrawing()
}
