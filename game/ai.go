package game

type Action struct {
}

func (g *Game) UpdateAi() {
	g.AiCar.ApplyGravity()
	forward := g.AiCar.Forward()
	g.AiCar.CarPosition.AddScaledVector(forward, g.AiCar.Velocity)
}

func (g *Game) ChangeAiVelocity(speedGradient float32) {
	if speedGradient > 0.5 {
		speedGradient = 0.5
	}
	if speedGradient < -0.5 {
		speedGradient = -0.5
	}
	g.AiCar.Velocity += speedGradient
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
