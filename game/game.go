package game

import "war-game-poc/input"

func UpdateGame(
	keyboardState input.KeyboardState,
	mouseState input.MouseState,
	player *Player,
) {
	if keyboardState.MoveFront {
		player.PlayerPosition.X += player.HeadingUnitVector.X
		player.PlayerPosition.Y += player.HeadingUnitVector.Y
		player.PlayerPosition.Z += player.HeadingUnitVector.Z
	}

	if keyboardState.MoveBack {
		player.PlayerPosition.X -= player.HeadingUnitVector.X
		player.PlayerPosition.Y -= player.HeadingUnitVector.Y
		player.PlayerPosition.Z -= player.HeadingUnitVector.Z

	}
}
