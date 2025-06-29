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
		0,
		1,
		0.1,
	)

	return &Player{
		Health:            health,
		PlayerPosition:    playerPosition,
		HeadingUnitVector: headingUnitVector,
	}
}

const wheelOffset = 5.0
const floorOffset = 2.5

func (p *Player) GetFrontWheelPosition() Position {
	forward := p.HeadingUnitVector
	worldUp := UnitVector{X: 0, Y: 1, Z: 0}
	right := Cross(forward, worldUp)
	localUp := Cross(right, forward)
	localDown := UnitVector{
		X: -localUp.X,
		Y: -localUp.Y,
		Z: -localUp.Z,
	}
	return Position{
		X: p.PlayerPosition.X + localDown.X*wheelOffset + forward.X*floorOffset,
		Y: p.PlayerPosition.Y + localDown.Y*wheelOffset + forward.Y*floorOffset,
		Z: p.PlayerPosition.Z + localDown.Z*wheelOffset + forward.Z*floorOffset,
	}
}

func (p *Player) GetRearWheelPosition() Position {
	forward := p.HeadingUnitVector
	worldUp := UnitVector{X: 0, Y: 1, Z: 0}
	right := Cross(forward, worldUp)
	localUp := Cross(right, forward)
	localDown := UnitVector{
		X: -localUp.X,
		Y: -localUp.Y,
		Z: -localUp.Z,
	}
	return Position{
		X: p.PlayerPosition.X + localDown.X*wheelOffset - forward.X*floorOffset,
		Y: p.PlayerPosition.Y + localDown.Y*wheelOffset - forward.Y*floorOffset,
		Z: p.PlayerPosition.Z + localDown.Z*wheelOffset - forward.Z*floorOffset,
	}
}
