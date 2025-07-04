package main

import (
	"fmt"
	ort "github.com/yalue/onnxruntime_go"
	"log"
)

func main() {
	/*output.InitWindow(1600, 900)
	defer output.CloseWindow()

	var player = game.CreatePlayer()

	for !output.ShouldClose() {

		// gather input
		keyboardInput := input.GetKeyboardInput()
		mouseInput := input.GetMouseInput()
		fmt.Println(keyboardInput)
		fmt.Println(mouseInput)

		// update game
		game.UpdateGame(
			keyboardInput,
			mouseInput,
			player,
		)

		// draw output
		output.DrawOutput(
			player,
		)

	}*/
	// hello
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
