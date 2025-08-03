package game

type Car struct {
	Health                 int32
	CarPosition            Position
	Yaw                    float32
	Pitch                  float32
	Roll                   float32
	Velocity               float32
	TargetVelocityGradient float32
	TargetRotationGradient float32
}

func CreateCar(
	position Position,
) *Car {
	var health int32 = 100
	return &Car{
		Health:                 health,
		CarPosition:            position,
		Yaw:                    0,
		Pitch:                  0,
		Roll:                   0,
		Velocity:               0,
		TargetVelocityGradient: 0,
		TargetRotationGradient: 0,
	}
}

const wheelOffset = 0.3
const floorOffset = 2.5

func (car *Car) GetFrontWheelPosition() Position {
	forward := car.Forward()
	worldUp := UnitVector{X: 0, Y: 1, Z: 0}
	right := Cross(forward, worldUp)
	localUp := Cross(right, forward)
	localDown := UnitVector{
		X: -localUp.X,
		Y: -localUp.Y,
		Z: -localUp.Z,
	}
	return Position{
		X: car.CarPosition.X + localDown.X*wheelOffset + forward.X*floorOffset,
		Y: car.CarPosition.Y + localDown.Y*wheelOffset + forward.Y*floorOffset,
		Z: car.CarPosition.Z + localDown.Z*wheelOffset + forward.Z*floorOffset,
	}
}

func (car *Car) GetRearWheelPosition() Position {
	forward := car.Forward()
	worldUp := UnitVector{X: 0, Y: 1, Z: 0}
	right := Cross(forward, worldUp)
	localUp := Cross(right, forward)
	localDown := UnitVector{
		X: -localUp.X,
		Y: -localUp.Y,
		Z: -localUp.Z,
	}
	return Position{
		X: car.CarPosition.X + localDown.X*wheelOffset - forward.X*floorOffset,
		Y: car.CarPosition.Y + localDown.Y*wheelOffset - forward.Y*floorOffset,
		Z: car.CarPosition.Z + localDown.Z*wheelOffset - forward.Z*floorOffset,
	}
}
