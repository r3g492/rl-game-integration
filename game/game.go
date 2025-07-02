package game

import "war-game-poc/input"

func UpdateGame(
	keyboardState input.KeyboardState,
	mouseState input.MouseState,
	player *Player,
) {
	forward := player.Forward()
	moveSpeed := float32(0.2)
	if keyboardState.MoveFront {
		player.PlayerPosition.X += forward.X * moveSpeed
		player.PlayerPosition.Y += forward.Y * moveSpeed
		player.PlayerPosition.Z += forward.Z * moveSpeed
	}
	if keyboardState.MoveBack {
		player.PlayerPosition.X -= forward.X * moveSpeed
		player.PlayerPosition.Y -= forward.Y * moveSpeed
		player.PlayerPosition.Z -= forward.Z * moveSpeed
	}
	if keyboardState.MoveRight {
		player.Yaw -= 0.1
	}
	if keyboardState.MoveLeft {
		player.Yaw += 0.1
	}

	if keyboardState.Jump {
		player.PlayerPosition.Y += 1
	}

	// way to implement gravity
	if player.GetFrontWheelPosition().Y <= 0 {

	}

	if player.GetRearWheelPosition().Y <= 0 {

	}

	if player.PlayerPosition.Y > 0 {
		player.PlayerPosition.Y -= 0.1
	}
}
