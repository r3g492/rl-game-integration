package main

import (
	"encoding/json"
	"fmt"
	ort "github.com/yalue/onnxruntime_go"
	"log"
	"net/http"
	"war-game-poc/game"
	"war-game-poc/input"
	"war-game-poc/output"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "gunwoo")
}

type StepResponse struct {
	Observation int  `json:"observation"`
	Reward      int  `json:"reward"`
	Done        bool `json:"done"`
}

// POST /reset
func resetHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	resp := map[string]int{"observation": 5}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("reset ")
	json.NewEncoder(w).Encode(resp)
}

// POST /step
func stepHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req struct {
		Action int `json:"action"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	resp := StepResponse{
		Observation: 2,
		Reward:      1,
		Done:        true,
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("step ", req.Action)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// init http part
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/step", stepHandler)
	fmt.Println("Server running on http://localhost:8080")
	go func() {
		fmt.Println("Server running on http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// do game part
	output.InitWindow(1600, 900)
	defer output.CloseWindow()

	var g = game.NewGame()

	for !output.ShouldClose() {

		// gather input
		keyboardInput := input.GetKeyboardInput()
		_ = input.GetMouseInput()

		// update game
		g.UpdatePlayer(keyboardInput)
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
