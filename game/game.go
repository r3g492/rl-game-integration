package game

import (
	"math/rand"
	"war-game-poc/input"
)

type State struct {
	Rps int // 0: Rock 1: paper 2: scissor
}

var gameState State
var PlayerCar *Car
var AiCar *Car

func GetRps() int {
	return gameState.Rps
}

func Duel(input int) int {
	opponent := gameState.Rps
	var reward = 0
	// 0: Rock, 1: Paper, 2: Scissors
	if input == opponent {
		reward = 0
	} else if (input == 0 && opponent == 2) || // Rock beats Scissors
		(input == 1 && opponent == 0) || // Paper beats Rock
		(input == 2 && opponent == 1) { // Scissors beats Paper
		reward = 1
	} else {
		reward = -1
	}
	gameState.Rps = rand.Intn(3)
	return reward
}

func UpdateGame(
	keyboardState input.KeyboardState,
) (*Car, *Car) {
	forward := PlayerCar.Forward()
	moveSpeed := float32(0.2)
	if keyboardState.MoveFront {
		PlayerCar.CarPosition.X += forward.X * moveSpeed
		PlayerCar.CarPosition.Y += forward.Y * moveSpeed
		PlayerCar.CarPosition.Z += forward.Z * moveSpeed
	}
	if keyboardState.MoveBack {
		PlayerCar.CarPosition.X -= forward.X * moveSpeed
		PlayerCar.CarPosition.Y -= forward.Y * moveSpeed
		PlayerCar.CarPosition.Z -= forward.Z * moveSpeed
	}
	if keyboardState.MoveRight {
		PlayerCar.Yaw -= 0.1
	}
	if keyboardState.MoveLeft {
		PlayerCar.Yaw += 0.1
	}

	if keyboardState.Jump {
		PlayerCar.CarPosition.Y += 1
	}

	if keyboardState.Reset {
		ResetGame()
	}

	// way to implement gravity
	if PlayerCar.GetFrontWheelPosition().Y <= 0 {

	}

	if PlayerCar.GetRearWheelPosition().Y <= 0 {

	}

	if PlayerCar.CarPosition.Y > 0 {
		PlayerCar.CarPosition.Y -= 0.1
	}

	if AiCar.CarPosition.Y > 0 {
		AiCar.CarPosition.Y -= 0.1
	}

	return PlayerCar, AiCar
}

func ResetGame() {
	PlayerCar = CreateCar(
		Position{
			X: 0,
			Y: 5,
			Z: 0,
		},
	)
	AiCar = CreateCar(
		Position{
			X: 10,
			Y: 3,
			Z: 10,
		},
	)
}
