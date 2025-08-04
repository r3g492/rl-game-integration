package game

import "war-game-poc/utility"

func (g *Game) UpdateAi(
	dt float32,
) {
	g.AiCar.ApplyGravity()
	forward := g.AiCar.Forward()

	// velocity
	g.AiCar.Velocity = utility.Suppress(
		g.AiCar.Velocity,
		-0.001,
		0.001,
	)
	g.AiCar.CarPosition.AddScaledVector(forward, g.AiCar.Velocity)
	g.AiCar.TargetVelocityGradient = utility.Clamp(
		g.AiCar.TargetVelocityGradient,
		-1.0,
		1.0,
	)
	g.AiCar.Velocity = utility.Friction(
		g.AiCar.Velocity,
		0.05,
	)
	g.AiCar.Velocity += g.AiCar.TargetVelocityGradient * dt * 5
	g.AiCar.Velocity = utility.Clamp(
		g.AiCar.Velocity,
		MinVelocity,
		MaxVelocity,
	)

	// rotation
	g.AiCar.TargetRotationGradient = utility.Clamp(
		g.AiCar.TargetRotationGradient,
		-1.0,
		1.0,
	)
	g.AiCar.TargetRotationGradient = utility.Suppress(
		g.AiCar.TargetRotationGradient,
		-0.5,
		0.5,
	)

	if g.AiCar.Velocity >= 0 {
		g.AiCar.Yaw += g.AiCar.TargetRotationGradient * dt * 2
	} else {
		g.AiCar.Yaw -= g.AiCar.TargetRotationGradient * dt * 2
	}
}

func (g *Game) friction(
	dt float32,
) {
	if g.AiCar.Velocity > 0 {
		g.AiCar.Velocity -= dt
		if g.AiCar.Velocity < 0 {
			g.AiCar.Velocity = 0
		}
	}
	if g.AiCar.Velocity < 0 {
		g.AiCar.Velocity += dt
		if g.AiCar.Velocity > MaxVelocity {
			g.AiCar.Velocity = MaxVelocity
		}
	}
}

var (
	MaxVelocity float32 = 2
	MinVelocity float32 = -2
)

func (g *Game) ChangeAiTargetVelocity(speedGradient float32) {
	// fmt.Println("change AiTargetVelocity: ", speedGradient)
	g.AiCar.TargetVelocityGradient = speedGradient
}

func (g *Game) ChangeAiTargetRotation(rotationGradient float32) {
	// fmt.Println("change AiTargetRotation: ", rotationGradient)
	g.AiCar.TargetRotationGradient = rotationGradient
}

func (g *Game) SaveAiPrevPosition(prev Position) {

}
