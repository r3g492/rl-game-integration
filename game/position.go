package game

type Position struct {
	X float32
	Y float32
	Z float32
}

func Distance(
	pos1 Position,
	pos2 Position,
) float32 {
	dx := pos1.X - pos2.X
	dy := pos1.Y - pos2.Y
	dz := pos1.Z - pos2.Z
	distance := dx*dx + dy*dy + dz*dz
	return distance
}
