package pixeldata

import (
	"fmt"
	"math/rand"
	"time"
)

// PixelJSON is the Go equivalent of the C# PixelJSON class.
type PixelJSON struct {
	PixelID     string `json:"pixelID"`
	PixelDataXY [][]int `json:"pixelDataXY"`
	PixelLabel  string `json:"pixelLabel"`
}

// GeneratePixelData creates a new PixelJSON object with a randomly generated pattern.
func GeneratePixelData() PixelJSON {
	pixelDrawXY := make([][]int, 16)
	for i := range pixelDrawXY {
		pixelDrawXY[i] = make([]int, 16)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	z := r.Intn(16)
	var d, sd int

	for y := 0; y <= 15; y++ {
		for x := 0; x <= 15; x++ {
			if x == z {
				pixelDrawXY[y][x] = 1
			} else {
				pixelDrawXY[y][x] = 0
			}
		}

		d = r.Intn(3) - 1 // -1, 0, 1
		if y != 0 {
			if d == 1 && sd == -1 {
				d = r.Intn(2) - 1 // -1, 0
			} else if d == -1 && sd == 1 {
				d = r.Intn(2) // 0, 1
			}
		}
		sd = d
		z += d

		if z == -1 {
			z++
		}
		if z == 16 {
			z--
		}
	}

	// Thicken the line
	for y := 0; y < 15; y += 2 {
		for x := 0; x < 15; x += 2 {
			if pixelDrawXY[y][x] == 1 || pixelDrawXY[y+1][x+1] == 1 || pixelDrawXY[y+1][x] == 1 || pixelDrawXY[y][x+1] == 1 {
				pixelDrawXY[y][x] = 1
				pixelDrawXY[y+1][x+1] = 1
				pixelDrawXY[y+1][x] = 1
				pixelDrawXY[y][x+1] = 1
			}
		}
	}

	return PixelJSON{
		PixelID:     hexGen(),
		PixelLabel:  "-",
		PixelDataXY: pixelDrawXY,
	}
}

// hexGen generates a random hexadecimal string.
func hexGen() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(999999999-10000000) + 10000000
	return fmt.Sprintf("%X", num)
}
