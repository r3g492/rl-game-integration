package game

import (
	"math/rand"
	"war-game-poc/input"
)

type State struct {
	Rps int // 0: Rock 1: paper 2: scissor
}

var gameState State

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
	player *Car,
) {
	forward := player.Forward()
	moveSpeed := float32(0.2)
	if keyboardState.MoveFront {
		player.CarPosition.X += forward.X * moveSpeed
		player.CarPosition.Y += forward.Y * moveSpeed
		player.CarPosition.Z += forward.Z * moveSpeed
	}
	if keyboardState.MoveBack {
		player.CarPosition.X -= forward.X * moveSpeed
		player.CarPosition.Y -= forward.Y * moveSpeed
		player.CarPosition.Z -= forward.Z * moveSpeed
	}
	if keyboardState.MoveRight {
		player.Yaw -= 0.1
	}
	if keyboardState.MoveLeft {
		player.Yaw += 0.1
	}

	if keyboardState.Jump {
		player.CarPosition.Y += 1
	}

	// way to implement gravity
	if player.GetFrontWheelPosition().Y <= 0 {

	}

	if player.GetRearWheelPosition().Y <= 0 {

	}

	if player.CarPosition.Y > 0 {
		player.CarPosition.Y -= 0.1
	}
}
