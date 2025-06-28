package output

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"war-game-poc/game"
)

func DrawOutput(
	player *game.Player,
) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.BeginMode3D(
		rl.NewCamera3D(
			rl.Vector3{X: 10, Y: 10, Z: 10},
			rl.Vector3{X: 0, Y: 0, Z: 0},
			rl.Vector3{X: 0, Y: 1, Z: 0},
			90,
			rl.CameraPerspective,
		),
	)

	rl.DrawSphere(
		rl.Vector3{X: player.PositionX, Y: player.PositionY, Z: player.PositionZ},
		1,
		rl.Red,
	)

	rl.DrawGrid(100, 1)
	rl.EndMode3D()
	rl.EndDrawing()
}
