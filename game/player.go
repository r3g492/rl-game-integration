package game

type Player struct {
	Health         int32
	PlayerPosition Position
	Yaw            float32
	Pitch          float32
	Roll           float32
}

func CreatePlayer() *Player {
	var health int32 = 100
	var playerPosition = Position{
		X: 0,
		Y: 5,
		Z: 0,
	}

	return &Player{
		Health:         health,
		PlayerPosition: playerPosition,
		Yaw:            0,
		Pitch:          0,
		Roll:           0,
	}
}

const wheelOffset = 0.3
const floorOffset = 2.5

func (p *Player) GetFrontWheelPosition() Position {
	forward := p.Forward()
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
	forward := p.Forward()
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
