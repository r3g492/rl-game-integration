package game

import (
	"war-game-poc/input"
)

const (
	moveSpeed  = float32(0.2)
	turnSpeed  = 0.1
	gravity    = 0.1
	jumpHeight = 1.0
)

type Game struct {
	PlayerCar *Car
	AiCar     *Car
	Goal      Position
}

func (g *Game) UpdateAi() {
	g.AiCar.ApplyGravity()
}

func (g *Game) UpdatePlayer(keyboardState input.KeyboardState) {
	forward := g.PlayerCar.Forward()

	if keyboardState.MoveFront {
		g.PlayerCar.CarPosition.AddScaledVector(forward, moveSpeed)
	}
	if keyboardState.MoveBack {
		g.PlayerCar.CarPosition.AddScaledVector(forward, -moveSpeed)
	}
	if keyboardState.MoveRight {
		g.PlayerCar.Yaw -= turnSpeed
	}
	if keyboardState.MoveLeft {
		g.PlayerCar.Yaw += turnSpeed
	}
	if keyboardState.Jump && g.PlayerCar.CarPosition.Y <= 0 {
		g.PlayerCar.CarPosition.Y += jumpHeight
	}
	if keyboardState.Reset {
		g.Reset()
	}

	g.PlayerCar.ApplyGravity()
}

func (c *Car) ApplyGravity() {
	if c.CarPosition.Y > 0 {
		c.CarPosition.Y -= gravity
		if c.CarPosition.Y < 0 {
			c.CarPosition.Y = 0
		}
	}
}

func (p *Position) AddScaledVector(vec UnitVector, scale float32) {
	p.X += vec.X * scale
	p.Y += vec.Y * scale
	p.Z += vec.Z * scale
}

func (g *Game) Reset() {
	g.PlayerCar = CreateCar(Position{X: 0, Y: 5, Z: 0})
	g.AiCar = CreateCar(Position{X: 10, Y: 3, Z: 10})
	g.Goal = Position{X: 0, Y: 0, Z: 30}
}

func (g *Game) CheckGoalIn() bool {
	const goalThreshold = 1.0
	dx := g.AiCar.CarPosition.X - g.Goal.X
	dy := g.AiCar.CarPosition.Y - g.Goal.Y
	dz := g.AiCar.CarPosition.Z - g.Goal.Z
	distanceSq := dx*dx + dy*dy + dz*dz
	return distanceSq <= goalThreshold*goalThreshold
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}
