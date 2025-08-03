package output

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
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
	var wheelPosition1 = player.GetFrontWheelPosition()
	rl.PushMatrix()
	rl.Translatef(wheelPosition1.X, wheelPosition1.Y, wheelPosition1.Z)
	rl.Rotatef(90, player.Forward().X, player.Forward().Y, player.Forward().Z)
	rl.DrawCylinder(rl.Vector3{0, 0, 0}, 0.5, 0.5, 2, 60, rl.Green)
	rl.PopMatrix()

	rl.PushMatrix()
	rl.Translatef(wheelPosition1.X, wheelPosition1.Y, wheelPosition1.Z)
	rl.Rotatef(-90, player.Forward().X, player.Forward().Y, player.Forward().Z)
	rl.DrawCylinder(rl.Vector3{0, 0, 0}, 0.5, 0.5, 2, 60, rl.Green)
	rl.PopMatrix()

	var wheelPosition2 = player.GetRearWheelPosition()
	rl.PushMatrix()
	rl.Translatef(wheelPosition2.X, wheelPosition2.Y, wheelPosition2.Z)
	rl.Rotatef(90, player.Forward().X, player.Forward().Y, player.Forward().Z)
	rl.DrawCylinder(rl.Vector3{0, 0, 0}, 0.5, 0.5, 2, 60, rl.Green)
	rl.PopMatrix()

	rl.PushMatrix()
	rl.Translatef(wheelPosition2.X, wheelPosition2.Y, wheelPosition2.Z)
	rl.Rotatef(-90, player.Forward().X, player.Forward().Y, player.Forward().Z)
	rl.DrawCylinder(rl.Vector3{0, 0, 0}, 0.5, 0.5, 2, 60, rl.Green)
	rl.PopMatrix()

	yawRad := math.Atan2(float64(player.Forward().X), float64(player.Forward().Z))
	yawDeg := float32(yawRad * (180.0 / math.Pi))
	rl.PushMatrix()
	rl.Translatef(player.CarPosition.X, player.CarPosition.Y, player.CarPosition.Z)
	rl.Rotatef(yawDeg, 0, 1, 0)
	rl.DrawCube(
		rl.Vector3{X: 0, Y: 0, Z: 0},
		3.0,
		1.0,
		8.0,
		rl.DarkGreen,
	)
	rl.PopMatrix()

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
