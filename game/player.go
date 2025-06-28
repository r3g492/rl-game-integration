package game

type Player struct {
	Health            int32
	PlayerPosition    Position
	HeadingUnitVector UnitVector
}

func CreatePlayer() *Player {
	var health int32 = 100
	var playerPosition = Position{
		X: 0,
		Y: 0,
		Z: 0,
	}
	var headingUnitVector, _ = NewUnitVector(
		1,
		1,
		1,
	)

	return &Player{
		Health:            health,
		PlayerPosition:    playerPosition,
		HeadingUnitVector: headingUnitVector,
	}
}
