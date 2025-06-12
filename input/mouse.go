package input

import rl "github.com/gen2brain/raylib-go/raylib"

type MouseState struct {
	X, Y           int32
	DeltaX, DeltaY int32
	Left, Right    bool
	Middle         bool
	LeftPressed    bool
	LeftReleased   bool
	RightPressed   bool
	RightReleased  bool
	WheelMove      float32
}

func GetMouseInput() MouseState {
	x := rl.GetMouseX()
	y := rl.GetMouseY()
	dx := rl.GetMouseDelta().X
	dy := rl.GetMouseDelta().Y
	return MouseState{
		X:             x,
		Y:             y,
		DeltaX:        int32(dx),
		DeltaY:        int32(dy),
		Left:          rl.IsMouseButtonDown(rl.MouseLeftButton),
		Right:         rl.IsMouseButtonDown(rl.MouseRightButton),
		Middle:        rl.IsMouseButtonDown(rl.MouseMiddleButton),
		LeftPressed:   rl.IsMouseButtonPressed(rl.MouseLeftButton),
		LeftReleased:  rl.IsMouseButtonReleased(rl.MouseLeftButton),
		RightPressed:  rl.IsMouseButtonPressed(rl.MouseRightButton),
		RightReleased: rl.IsMouseButtonReleased(rl.MouseRightButton),
		WheelMove:     rl.GetMouseWheelMove(),
	}
}
