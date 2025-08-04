package utility

func Clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func Suppress(x, lower, upper float32) float32 {
	if x > lower && x < upper {
		return 0
	}
	return x
}

func Friction(x, delta float32) float32 {
	if x > delta {
		return x - delta
	} else if x < -delta {
		return x + delta
	}
	return 0
}

func Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}
