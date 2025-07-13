package game

import (
	"math"
)

type UnitVector struct {
	X, Y, Z float32
}

func Cross(a, b UnitVector) UnitVector {
	return UnitVector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func (p *Car) Forward() UnitVector {
	cosPitch := float32(math.Cos(float64(p.Pitch)))
	return UnitVector{
		X: float32(math.Sin(float64(p.Yaw))) * cosPitch,
		Y: float32(math.Sin(float64(p.Pitch))),
		Z: float32(math.Cos(float64(p.Yaw))) * cosPitch,
	}
}
