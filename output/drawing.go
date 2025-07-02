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
			rl.Vector3{
				X: player.PlayerPosition.X - 12*player.Forward().X,
				Y: player.PlayerPosition.Y + 12,
				Z: player.PlayerPosition.Z - 12*player.Forward().Z,
			},
			rl.Vector3{
				X: player.PlayerPosition.X + player.Forward().X,
				Y: player.PlayerPosition.Y + player.Forward().Y,
				Z: player.PlayerPosition.Z + player.Forward().Z,
			},
			rl.Vector3{X: 0, Y: 1, Z: 0},
			90,
			rl.CameraPerspective,
		),
	)

	rl.DrawSphere(
		rl.Vector3{X: player.PlayerPosition.X, Y: player.PlayerPosition.Y, Z: player.PlayerPosition.Z},
		1,
		rl.Red,
	)

	var wheelPosition1 = player.GetFrontWheelPosition()
	rl.DrawSphere(
		rl.Vector3{X: wheelPosition1.X, Y: wheelPosition1.Y, Z: wheelPosition1.Z},
		0.3,
		rl.Blue,
	)

	var wheelPosition2 = player.GetRearWheelPosition()
	rl.DrawSphere(
		rl.Vector3{X: wheelPosition2.X, Y: wheelPosition2.Y, Z: wheelPosition2.Z},
		0.3,
		rl.Blue,
	)

	const arrowLength float32 = 5
	start := rl.Vector3{
		X: player.PlayerPosition.X,
		Y: player.PlayerPosition.Y,
		Z: player.PlayerPosition.Z,
	}
	end := rl.Vector3{
		X: player.PlayerPosition.X + player.Forward().X*arrowLength,
		Y: player.PlayerPosition.Y + player.Forward().Y*arrowLength,
		Z: player.PlayerPosition.Z + player.Forward().Z*arrowLength,
	}
	rl.DrawLine3D(start, end, rl.Red)

	rl.DrawGrid(100, 1)
	rl.EndMode3D()
	rl.EndDrawing()
}
