package game

import (
	"fmt"
	"time"
	"war-game-poc/input"
	"war-game-poc/utility"
)

const (
	moveSpeed  = float32(0.2)
	turnSpeed  = 0.1
	gravity    = 0.1
	jumpHeight = 1.0
)

type Game struct {
	PlayerCar   *Car
	AiCar       *Car
	Goal        Position
	GoalReached bool
	Reward      float32
	StartTime   time.Time
}

func (g *Game) ControlOptions(keyboardState input.KeyboardState) {
	if keyboardState.Reset {
		g.Reset()
	}
}

func (g *Game) ControlPlayer(keyboardState input.KeyboardState) {
	if keyboardState.MoveFront {
		g.PlayerCar.TargetVelocityGradient = 1
	} else if keyboardState.MoveBack {
		g.PlayerCar.TargetVelocityGradient = -1
	} else {
		g.PlayerCar.TargetVelocityGradient = 0
	}
	if keyboardState.MoveRight {
		g.PlayerCar.TargetRotationGradient = -1
	} else if keyboardState.MoveLeft {
		g.PlayerCar.TargetRotationGradient = 1
	} else {
		g.PlayerCar.TargetRotationGradient = 0
	}

}

func (g *Game) UpdatePlayer(
	dt float32,
) {
	g.PlayerCar.ApplyGravity()
	forward := g.PlayerCar.Forward()

	// velocity
	g.PlayerCar.Velocity = utility.Suppress(
		g.PlayerCar.Velocity,
		-0.001,
		0.001,
	)
	g.PlayerCar.CarPosition.AddScaledVector(forward, g.PlayerCar.Velocity)
	g.PlayerCar.TargetVelocityGradient = utility.Clamp(
		g.PlayerCar.TargetVelocityGradient,
		-1.0,
		1.0,
	)
	g.PlayerCar.Velocity = utility.Friction(
		g.PlayerCar.Velocity,
		0.05,
	)
	g.PlayerCar.Velocity += g.PlayerCar.TargetVelocityGradient * dt * 5
	g.PlayerCar.Velocity = utility.Clamp(
		g.PlayerCar.Velocity,
		MinVelocity,
		MaxVelocity,
	)

	// rotation
	g.PlayerCar.TargetRotationGradient = utility.Clamp(
		g.PlayerCar.TargetRotationGradient,
		-1.0,
		1.0,
	)
	g.PlayerCar.TargetRotationGradient = utility.Suppress(
		g.PlayerCar.TargetRotationGradient,
		-0.5,
		0.5,
	)

	if g.PlayerCar.Velocity >= 0 {
		g.PlayerCar.Yaw += g.PlayerCar.TargetRotationGradient * dt * 2
	} else {
		g.PlayerCar.Yaw -= g.PlayerCar.TargetRotationGradient * dt * 2
	}
}

func (car *Car) ApplyGravity() {
	if car.CarPosition.Y > 0 {
		car.CarPosition.Y -= gravity
		if car.CarPosition.Y < 0 {
			car.CarPosition.Y = 0
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
	g.GoalReached = false
	g.Reward = 0
	g.StartTime = time.Now()
}

func (g *Game) CheckGoalIn() bool {
	const goalThreshold = 1.5
	dx := g.AiCar.CarPosition.X - g.Goal.X
	dy := g.AiCar.CarPosition.Y - g.Goal.Y
	dz := g.AiCar.CarPosition.Z - g.Goal.Z
	distanceSq := dx*dx + dy*dy + dz*dz
	return distanceSq <= goalThreshold*goalThreshold
}

func (g *Game) CheckGoalOut() bool {
	const goalOutThreshold = 30.0
	dx := g.AiCar.CarPosition.X - g.Goal.X
	dy := g.AiCar.CarPosition.Y - g.Goal.Y
	dz := g.AiCar.CarPosition.Z - g.Goal.Z
	distanceSq := dx*dx + dy*dy + dz*dz
	timeElapsed := float32(time.Since(g.StartTime).Seconds())
	if timeElapsed > 10 {
		return true
	}
	return distanceSq > goalOutThreshold*goalOutThreshold
}

func (g *Game) DistanceFromGoal() float32 {
	dx := g.AiCar.CarPosition.X - g.Goal.X
	dy := g.AiCar.CarPosition.Y - g.Goal.Y
	dz := g.AiCar.CarPosition.Z - g.Goal.Z
	distance := dx*dx + dy*dy + dz*dz
	return distance
}

func NewGame() *Game {
	g := &Game{}
	g.Reset()
	return g
}

func (g *Game) Won() bool {
	return g.GoalReached && g.Reward > 0
}

func (g *Game) Lost() bool {
	return g.GoalReached && g.Reward < 0
}

func (g *Game) GoalUpdate() {
	if g.GoalReached {
		return
	}
	if g.CheckGoalIn() {
		g.Reward = 3
		g.GoalReached = true
		fmt.Println("goal in, reward: ", g.Reward)
	}
	if g.CheckGoalOut() {
		g.Reward = -g.DistanceFromGoal() / 100
		g.GoalReached = true

		fmt.Println("goal out, reward: ", g.Reward)
	}
}
