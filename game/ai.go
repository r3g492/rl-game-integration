package game

type Action struct {
}

func (g *Game) UpdateAi() {
	g.AiCar.ApplyGravity()
	forward := g.AiCar.Forward()
	g.AiCar.CarPosition.AddScaledVector(forward, g.AiCar.Velocity)

	if g.AiCar.TargetVelocityGradient > 0.5 {
		g.AiCar.TargetVelocityGradient = 0.5
	}
	if g.AiCar.TargetVelocityGradient < -0.5 {
		g.AiCar.TargetVelocityGradient = -0.5
	}
	g.AiCar.Velocity += g.AiCar.TargetVelocityGradient
	if g.AiCar.Velocity > MaxVelocity {
		g.AiCar.Velocity = MaxVelocity
	}

	if g.AiCar.Velocity < MinVelocity {
		g.AiCar.Velocity = MinVelocity
	}

	if g.AiCar.TargetRotationGradient > 0.1 {
		g.AiCar.TargetRotationGradient = 0.1
	}
	if g.AiCar.TargetRotationGradient < -0.1 {
		g.AiCar.TargetRotationGradient = -0.1
	}
	g.AiCar.Yaw += g.AiCar.TargetRotationGradient
}

var (
	MaxVelocity float32 = 2
	MinVelocity float32 = -2
)

func (g *Game) ChangeAiTargetVelocity(speedGradient float32) {
	g.AiCar.TargetVelocityGradient = speedGradient
}

func (g *Game) ChangeAiTargetRotation(rotationGradient float32) {
	g.AiCar.TargetRotationGradient = rotationGradient
}
