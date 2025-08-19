package pixelai

import (
	"fmt"
	"pixel-draw-go/pixeldata"
	"strconv"
	"strings"
)

// AI analyzes the given PixelJSON data to determine if the pattern is a straight line.
func AI(data pixeldata.PixelJSON) pixeldata.PixelJSON {
	const xyLength = 15
	var xy []string

	// Extract coordinates from the pixel data
	for y := 0; y < xyLength; y += 2 {
		for x := 0; x < xyLength; x += 2 {
			if data.PixelDataXY[y][x] == 1 {
				// Scale down the coordinates
				x1 := x / 2
				y1 := y / 2
				xy = append(xy, fmt.Sprintf("%d,%d", y1, x1))
			}
		}
	}

	// This was in the original code to see the coordinates
	for _, v := range xy {
		fmt.Println(v)
	}

	var downCount, upCount int

	// Analyze the change in y-coordinates
	for i := 1; i < len(xy); i++ {
		xyFirstParts := strings.Split(xy[i-1], ",")
		xySecondParts := strings.Split(xy[i], ",")

		yFirst, err1 := strconv.Atoi(xyFirstParts[0])
		ySecond, err2 := strconv.Atoi(xySecondParts[0])

		if err1 != nil || err2 != nil {
			// Handle error, maybe skip this pair
			continue
		}

		if yFirst < ySecond {
			downCount++
		}
		if yFirst > ySecond {
			upCount++
		}
	}

	// Check if the line is straight
	if upCount == 0 || downCount == 0 {
		fmt.Println("S")
	}

	return data
}

// ML is a placeholder for a machine learning model.
func ML(data pixeldata.PixelJSON) pixeldata.PixelJSON {
	return data
}
