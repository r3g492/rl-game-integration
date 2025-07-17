package output

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"war-game-poc/game"
)

func DrawGame(
	g *game.Game,
) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.BeginMode3D(
		rl.NewCamera3D(
			rl.Vector3{
				X: g.PlayerCar.CarPosition.X - 12*g.PlayerCar.Forward().X,
				Y: g.PlayerCar.CarPosition.Y + 12,
				Z: g.PlayerCar.CarPosition.Z - 12*g.PlayerCar.Forward().Z,
			},
			rl.Vector3{
				X: g.PlayerCar.CarPosition.X + g.PlayerCar.Forward().X,
				Y: g.PlayerCar.CarPosition.Y + g.PlayerCar.Forward().Y,
				Z: g.PlayerCar.CarPosition.Z + g.PlayerCar.Forward().Z,
			},
			rl.Vector3{X: 0, Y: 1, Z: 0},
			90,
			rl.CameraPerspective,
		),
	)

	drawCar(g.PlayerCar)
	drawCar(g.AiCar)

	rl.DrawSphere(
		rl.Vector3{X: g.Goal.X, Y: g.Goal.Y, Z: g.Goal.Z},
		1,
		rl.Blue,
	)

	rl.DrawGrid(100, 1)
	rl.EndMode3D()
	rl.EndDrawing()
}

func drawCar(player *game.Car) {
	rl.DrawSphere(
		rl.Vector3{X: player.CarPosition.X, Y: player.CarPosition.Y, Z: player.CarPosition.Z},
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
		X: player.CarPosition.X,
		Y: player.CarPosition.Y,
		Z: player.CarPosition.Z,
	}
	end := rl.Vector3{
		X: player.CarPosition.X + player.Forward().X*arrowLength,
		Y: player.CarPosition.Y + player.Forward().Y*arrowLength,
		Z: player.CarPosition.Z + player.Forward().Z*arrowLength,
	}
	rl.DrawLine3D(start, end, rl.Red)
}
