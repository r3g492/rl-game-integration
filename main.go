package main

import (
	"encoding/json"
	"fmt"
	ort "github.com/yalue/onnxruntime_go"
	"log"
	"net/http"
	"sync"
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

	var reward float32 = 0
	var done = g.IsDone()
	if done {
		reward = g.Reward
		g.Reset()
	}

	resp := train.StepResponse{
		Observation: obs,
		Reward:      reward,
		Done:        done,
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
	g.ChangeAiVelocity(req.SpeedGradient)
	g.ChangeAiRotation(req.RotationGradient)

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

	var reward float32 = 0
	var done = g.IsDone()
	if done {
		reward = g.AiCar.Velocity
	}

	resp := train.StepResponse{
		Observation: obs,
		Reward:      reward,
		Done:        done,
	}

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
	for !output.ShouldClose() {

		// gather input
		keyboardInput := input.GetKeyboardInput()
		_ = input.GetMouseInput()

		// update game
		g.ControlPlayer(keyboardInput)
		g.UpdatePlayer()
		g.UpdateAi()

		if g.CheckGoalIn() {
			g.Reset()
		}

		// draw output
		output.DrawGame(g)
	}

	// below code is for RPS. does not do anything at the moment.
	ort.SetSharedLibraryPath("./libonnxruntime/libonnxruntime.so")
	err := ort.InitializeEnvironment()
	if err != nil {
		panic(err)
	}
	defer ort.DestroyEnvironment()

	// Input: int32[1,1]
	inputData := []int32{1} // 0: Rock, 1: Paper, 2: Scissors
	inputShape := ort.NewShape(1, 1)
	inputTensor, err := ort.NewTensor(inputShape, inputData)
	if err != nil {
		log.Fatalf("failed to create input tensor: %v", err)
	}
	defer inputTensor.Destroy()

	// Output: int64[1]
	outputShape := ort.NewShape(1)
	outputTensor, err := ort.NewEmptyTensor[int64](outputShape)
	if err != nil {
		log.Fatalf("failed to create output tensor: %v", err)
	}
	defer outputTensor.Destroy()

	// Session (input/output names from Netron: "input" and "output")
	session, err := ort.NewAdvancedSession("./rps_agent_visible.onnx",
		[]string{"input"}, []string{"output"},
		[]ort.Value{inputTensor}, []ort.Value{outputTensor}, nil)
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	defer session.Destroy()

	// Run inference
	err = session.Run()
	if err != nil {
		log.Fatalf("failed to run session: %v", err)
	}

	// Get output values (predicted action)
	outputVals := outputTensor.GetData()
	fmt.Printf("Predicted action index: %d\n", outputVals[0])
	actions := []string{"Rock", "Paper", "Scissors"}
	if int(outputVals[0]) >= 0 && int(outputVals[0]) < len(actions) {
		fmt.Printf("Predicted action: %s\n", actions[outputVals[0]])
	} else {
		fmt.Printf("Invalid action index: %d\n", outputVals[0])
	}

}
