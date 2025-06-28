package game

import (
	"errors"
	"math"
)

type UnitVector struct {
	X, Y, Z float32
}

func NewUnitVector(x, y, z float32) (UnitVector, error) {
	sqrt := float32(math.Sqrt(float64(x*x + y*y + z*z)))
	if sqrt == 0 {
		return UnitVector{}, errors.New("cannot create unit vector from zero vector")
	}
	return UnitVector{X: x / sqrt, Y: y / sqrt, Z: z / sqrt}, nil
}

func Cross(a, b UnitVector) UnitVector {
	return UnitVector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
