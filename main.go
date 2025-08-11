package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"war-game-poc/game"
	"war-game-poc/input"
	"war-game-poc/output"
	"war-game-poc/train"
)

var (
	g     *game.Game
	gLock sync.Mutex
)

// POST /reset
func resetHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	gLock.Lock()
	defer gLock.Unlock()

	g.Reset()
	obs := train.Observation{
		CarX:     g.AiCar.CarPosition.X,
		CarY:     g.AiCar.CarPosition.Y,
		CarZ:     g.AiCar.CarPosition.Z,
		Velocity: g.AiCar.Velocity,
		Yaw:      g.AiCar.Yaw,
		GoalX:    g.Goal.X,
		GoalY:    g.Goal.Y,
		GoalZ:    g.Goal.Z,
	}

	resp := train.StepResponse{
		Observation: obs,
		Reward:      0,
		Done:        g.Done(),
		Truncated:   g.Truncated(),
		IsSuccess:   g.IsSuccess(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /step
func stepHandler(w http.ResponseWriter, r *http.Request) {
	gLock.Lock()
	defer gLock.Unlock()

	var req train.StepRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	g.ChangeAiTargetVelocity(req.SpeedGradient)
	g.ChangeAiTargetRotation(req.RotationGradient)

	obs := train.Observation{
		CarX:     g.AiCar.CarPosition.X,
		CarY:     g.AiCar.CarPosition.Y,
		CarZ:     g.AiCar.CarPosition.Z,
		Velocity: g.AiCar.Velocity,
		Yaw:      g.AiCar.Yaw,
		GoalX:    g.Goal.X,
		GoalY:    g.Goal.Y,
		GoalZ:    g.Goal.Z,
	}

	resp := train.StepResponse{
		Observation: obs,
		Reward:      g.Reward(),
		Done:        g.Done(),
		Truncated:   g.Truncated(),
		IsSuccess:   g.IsSuccess(),
	}

	g.SaveAiPrevPosition(
		game.Position{
			X: g.AiCar.CarPosition.X,
			Y: g.AiCar.CarPosition.Y,
			Z: g.AiCar.CarPosition.Z,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// init http part
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/step", stepHandler)
	go func() {
		fmt.Println("Server running on http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// do game part
	output.InitWindow(1600, 900)
	defer output.CloseWindow()

	g = game.NewGame()
	if g.AiCar == nil {
		panic("g.AiCar is nil after NewGame")
	}

	lastTime := time.Now()
	for !output.ShouldClose() {
		now := time.Now()
		dt := float32(now.Sub(lastTime).Seconds())
		lastTime = now
		// gather input
		keyboardInput := input.GetKeyboardInput()
		_ = input.GetMouseInput()
		g.ControlOptions(keyboardInput)
		// update game
		if !g.Done() {
			g.ControlPlayer(keyboardInput)
			g.UpdatePlayer(dt)
			g.UpdateAi(dt)
		}

		// draw output
		output.DrawGame(g)
	}
}
