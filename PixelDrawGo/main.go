package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"pixel-draw-go/jsondb"
	"pixel-draw-go/pixelai"
	"pixel-draw-go/pixeldata"
)

const k = 5 // k value for the k-NN algorithm

func main() {
	// Define a flag for test run
	testRun := flag.Bool("testrun", false, "Run in non-interactive test mode")
	flag.Parse()

	if *testRun {
		runTest()
	} else {
		reader := bufio.NewReader(os.Stdin)
		runInteractive(reader)
	}
}

func runTest() {
	fmt.Println("--- Running Save/Load Test ---")
	testSaveAndLoad()
	fmt.Println("\n--- Running k-NN Prediction Test ---")
	testKNN()
	fmt.Println("\nAll tests completed successfully.")
}

func testSaveAndLoad() {
	// This test is designed to be run on a clean slate.
	// Let's remove the old test file if it exists.
	os.Remove("PixelTrainingData.jsonl")

	fmt.Println("Running in test mode...")
	pJSON := pixeldata.GeneratePixelData()
	fmt.Println("Generated Pixel ID:", pJSON.PixelID)
	pJSON.PixelLabel = "straight-test"
	fmt.Println("Labeled as:", pJSON.PixelLabel)

	if err := jsondb.PutData(pJSON); err != nil {
		log.Fatalf("Error saving data in test mode: %v", err)
	}
	fmt.Println("Data saved successfully.")

	data, err := jsondb.GetData()
	if err != nil {
		log.Fatalf("Error reading data in test mode: %v", err)
	}

	if len(data) > 0 && data[len(data)-1].PixelID == pJSON.PixelID {
		fmt.Println("Successfully read data back from the file.")
		fmt.Println("Save/Load Test completed successfully.")
	} else {
		log.Fatalf("Failed to read data back from the file or data mismatch.")
	}
}

func testKNN() {
	fmt.Println("Creating a dummy training set for k-NN test...")
	trainingData := []pixeldata.PixelJSON{}

	// Create 5 identical "straight" patterns
	for i := 0; i < 5; i++ {
		straightLine := make([][]int, 16)
		for r := range straightLine {
			straightLine[r] = make([]int, 16)
			if r == 8 { // A horizontal line in the middle
				for c := range straightLine[r] {
					straightLine[r][c] = 1
				}
			}
		}
		trainingData = append(trainingData, pixeldata.PixelJSON{PixelLabel: "straight", PixelDataXY: straightLine})
	}

	// Create a test pattern that is identical to the "straight" ones
	testPattern := pixeldata.PixelJSON{
		PixelDataXY: trainingData[0].PixelDataXY,
	}

	fmt.Println("Predicting the label of a known 'straight' pattern...")
	prediction := pixelai.Predict(trainingData, testPattern, k)

	fmt.Println("Predicted label:", prediction)
	if prediction == "straight" {
		fmt.Println("k-NN Test passed: Correctly predicted 'straight'.")
	} else {
		log.Fatalf("k-NN Test failed: Expected 'straight', but got '%s'", prediction)
	}
}


func runInteractive(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Label a new pattern")
		fmt.Println("4. Predict a new pattern")
		fmt.Println("9. Exit")
		fmt.Print("# ")

		input, _ := reader.ReadString('\n')
		answer, _ := strconv.Atoi(strings.TrimSpace(input))

		switch answer {
		case 1:
			runLabeling(reader)
		case 4:
			runPrediction(reader)
		case 9:
			fmt.Println("Exiting application.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func runLabeling(reader *bufio.Reader) {
	pJSON := pixeldata.GeneratePixelData()
	pixelConsolAnalyze(pJSON)

	fmt.Println()
	fmt.Println("What is the line?")
	fmt.Println("1. straight")
	fmt.Println("2. curved")
	fmt.Println("3. sinewave")
	fmt.Print("# ")

	input, _ := reader.ReadString('\n')
	answer, _ := strconv.Atoi(strings.TrimSpace(input))

	switch answer {
	case 1:
		pJSON.PixelLabel = "straight"
	case 2:
		pJSON.PixelLabel = "curved"
	case 3:
		pJSON.PixelLabel = "sinewave"
	default:
		fmt.Println("Invalid choice. Data not saved.")
		return
	}

	if err := jsondb.PutData(pJSON); err != nil {
		log.Println("Error saving data:", err)
	} else {
		fmt.Println("Data saved.")
	}
}

func runPrediction(reader *bufio.Reader) {
	fmt.Println("\n--- Prediction Mode ---")
	trainingData, err := jsondb.GetData()
	if err != nil {
		log.Println("Error getting training data:", err)
		return
	}

	if len(trainingData) < k {
		fmt.Printf("Not enough training data. Need at least %d labeled patterns. You have %d.\n", k, len(trainingData))
		return
	}

	fmt.Printf("Loaded %d labeled patterns for prediction.\n", len(trainingData))
	newPattern := pixeldata.GeneratePixelData()

	prediction := pixelai.Predict(trainingData, newPattern, k)

	newPattern.PixelLabel = "PREDICTED: " + prediction
	pixelConsolAnalyze(newPattern)

	fmt.Println("\nPress Enter to continue...")
	reader.ReadString('\n')
}

func pixelConsolAnalyze(data pixeldata.PixelJSON) {
	fmt.Println("──────────────────────────────")
	fmt.Println()
	fmt.Printf("PixelID:    %s\n", data.PixelID)
	fmt.Printf("PixelLabel: %s\n", data.PixelLabel)
	fmt.Println()

	// Top border
	fmt.Print("┌")
	for i := 0; i < 23; i++ {
		switch i {
		case 2, 5, 8, 11, 14, 17, 20:
			fmt.Print("┬")
		case 9:
			fmt.Print("Y")
		case 10:
			fmt.Print("▼")
		case 12:
			fmt.Print("X")
		case 13:
			fmt.Print("►")
		default:
			fmt.Print("─")
		}
	}
	fmt.Println("┐")

	// Middle grid
	for y := 0; y < 16; y++ {
		if y == 2 || y == 4 || y == 6 || y == 8 || y == 10 || y == 12 || y == 14 {
			fmt.Println("├──┼──┼──┼──┼──┼──┼──┼──┤")
		}

		fmt.Print("│")
		for x := 0; x < 16; x++ {
			if data.PixelDataXY[y][x] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("█")
			}
			if x == 1 || x == 3 || x == 5 || x == 7 || x == 9 || x == 11 || x == 13 {
				fmt.Print("│")
			}
		}
		fmt.Println("│")
	}

	// Bottom border
	fmt.Print("└")
	for i := 0; i < 23; i++ {
		if i == 2 || i == 5 || i == 8 || i == 11 || i == 14 || i == 17 || i == 20 {
			fmt.Print("┴")
		} else {
			fmt.Print("─")
		}
	}
	fmt.Println("┘")
	fmt.Println()
}
