package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	moveFrontKey int32 = rl.KeyW
	moveLeftKey  int32 = rl.KeyA
	moveBackKey  int32 = rl.KeyS
	moveRightKey int32 = rl.KeyD
	useKey       int32 = rl.KeyE
	jumpKey      int32 = rl.KeySpace
	reset        int32 = rl.KeyR
)

type KeyboardState struct {
	MoveFront, MoveBack, MoveLeft, MoveRight bool
	Use, Jump                                bool
	Reset                                    bool
}

func GetKeyboardInput() KeyboardState {
	return KeyboardState{
		MoveFront: rl.IsKeyDown(moveFrontKey),
		MoveBack:  rl.IsKeyDown(moveBackKey),
		MoveLeft:  rl.IsKeyDown(moveLeftKey),
		MoveRight: rl.IsKeyDown(moveRightKey),
		Use:       rl.IsKeyDown(useKey),
		Jump:      rl.IsKeyDown(jumpKey),
		Reset:     rl.IsKeyDown(reset),
	}
}
