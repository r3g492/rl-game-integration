package game

type Action struct {
}

func (g *Game) UpdateAi() {
	g.AiCar.ApplyGravity()
	forward := g.AiCar.Forward()
	g.AiCar.CarPosition.AddScaledVector(forward, g.AiCar.Velocity)
}

var (
	MaxVelocity float32 = 2
	MinVelocity float32 = -2
)

func (g *Game) ChangeAiVelocity(speedGradient float32) {
	if speedGradient > 0.5 {
		speedGradient = 0.5
	}
	if speedGradient < -1.5 {
		speedGradient = -1.5
	}
	g.AiCar.Velocity += speedGradient
	if g.AiCar.Velocity > MaxVelocity {
		g.AiCar.Velocity = MaxVelocity
	}

	if g.AiCar.Velocity < MinVelocity {
		g.AiCar.Velocity = MinVelocity
	}
}

func (g *Game) ChangeAiRotation(rotationGradient float32) {
	if rotationGradient > 0.5 {
		rotationGradient = 0.5
	}
	if rotationGradient < -0.5 {
		rotationGradient = -0.5
	}
	g.AiCar.Yaw += rotationGradient
}
