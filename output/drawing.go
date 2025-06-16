package output

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawOutput() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	rl.BeginMode3D(
		rl.NewCamera3D(
			rl.Vector3{X: 10, Y: 10, Z: 10},
			rl.Vector3{X: 0, Y: 0, Z: 0},
			rl.Vector3{X: 0, Y: 1, Z: 0},
			90,
			rl.CameraPerspective,
		),
	)

	rl.DrawGrid(100, 1)
	rl.EndDrawing()
}
