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
	"pixel-draw-go/pixeldata"
)

func main() {
	// Define a flag for test run
	testRun := flag.Bool("testrun", false, "Run in non-interactive test mode")
	flag.Parse()

	if *testRun {
		runTest()
	} else {
		runInteractive()
	}
}

func runTest() {
	fmt.Println("Running in test mode...")
	// 1. Generate data
	pJSON := pixeldata.GeneratePixelData()
	fmt.Println("Generated Pixel ID:", pJSON.PixelID)

	// 2. Label it
	pJSON.PixelLabel = "straight-test"
	fmt.Println("Labeled as:", pJSON.PixelLabel)

	// 3. Save it
	if err := jsondb.PutData(pJSON); err != nil {
		log.Fatalf("Error saving data in test mode: %v", err)
	}
	fmt.Println("Data saved successfully.")

	// 4. Verify by reading
	data, err := jsondb.GetData()
	if err != nil {
		log.Fatalf("Error reading data in test mode: %v", err)
	}

	if len(data) > 0 && data[len(data)-1].PixelID == pJSON.PixelID {
		fmt.Println("Successfully read data back from the file.")
		fmt.Println("Test completed successfully.")
	} else {
		log.Fatalf("Failed to read data back from the file or data mismatch.")
	}
}

func runInteractive() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Generate new pixel data
		pJSON := pixeldata.GeneratePixelData()

		// Display the data in the console
		pixelConsolAnalyze(pJSON)

		// Prompt user for input
		fmt.Println()
		fmt.Println("What is the line?")
		fmt.Println("1. straight")
		fmt.Println("2. curved")
		fmt.Println("3. sinewave")
		fmt.Println("9. Exit")
		fmt.Print("# ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}

		// Trim whitespace and convert to integer
		answer, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Println("Invalid input, please enter a number.")
			continue
		}

		// Process user's choice
		if answer == 9 {
			break // Exit the loop
		}

		switch answer {
		case 1:
			pJSON.PixelLabel = "straight"
		case 2:
			pJSON.PixelLabel = "curved"
		case 3:
			pJSON.PixelLabel = "sinewave"
		default:
			fmt.Println("Invalid choice. Data not saved.")
			continue // Skip saving
		}

		// Save the labeled data
		if err := jsondb.PutData(pJSON); err != nil {
			log.Println("Error saving data:", err)
		} else {
			fmt.Println("Data saved.")
		}
	}

	fmt.Println("Exiting application.")
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
