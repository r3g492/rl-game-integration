package game

type Action struct {
}

func (g *Game) UpdateAi() {
	g.AiCar.ApplyGravity()
	forward := g.AiCar.Forward()
	g.AiCar.CarPosition.AddScaledVector(forward, g.AiCar.Velocity)
}

func (g *Game) ChangeAiVelocity(velocity float32) {
	g.AiCar.Velocity = velocity
}
