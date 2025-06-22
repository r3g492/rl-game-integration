package game

type Player struct {
	Health    int32
	PositionX float32
	PositionY float32
	PositionZ float32
}

func CreatePlayer() *Player {
	return &Player{
		Health:    100,
		PositionX: 0.0,
		PositionY: 0.0,
		PositionZ: 0.0,
	}
}
