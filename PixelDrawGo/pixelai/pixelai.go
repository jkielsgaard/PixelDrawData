package pixelai

import (
	"math"
	"pixel-draw-go/pixeldata"
	"sort"
)

// A struct to hold a neighbor's distance and label
type neighbor struct {
	distance float64
	label    string
}

// flatten converts a 2D grid into a 1D slice.
func flatten(grid [][]int) []int {
	flat := []int{}
	for _, row := range grid {
		flat = append(flat, row...)
	}
	return flat
}

// euclideanDistance calculates the Euclidean distance between two vectors.
func euclideanDistance(a, b []int) float64 {
	var sum float64
	for i := range a {
		sum += math.Pow(float64(a[i]-b[i]), 2)
	}
	return math.Sqrt(sum)
}

// Predict uses the k-NN algorithm to predict the label of a new pattern.
func Predict(trainingData []pixeldata.PixelJSON, newPattern pixeldata.PixelJSON, k int) string {
	if len(trainingData) < k {
		return "Not enough training data"
	}

	flatNewPattern := flatten(newPattern.PixelDataXY)
	var neighbors []neighbor

	// Calculate the distance to each training data point
	for _, item := range trainingData {
		flatItem := flatten(item.PixelDataXY)
		dist := euclideanDistance(flatNewPattern, flatItem)
		neighbors = append(neighbors, neighbor{distance: dist, label: item.PixelLabel})
	}

	// Sort neighbors by distance
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].distance < neighbors[j].distance
	})

	// Get the top k neighbors
	kNeighbors := neighbors[:k]

	// Count the votes for each label
	votes := make(map[string]int)
	for _, n := range kNeighbors {
		votes[n.label]++
	}

	// Find the label with the most votes
	var maxVotes int
	var prediction string
	for label, count := range votes {
		if count > maxVotes {
			maxVotes = count
			prediction = label
		}
	}

	return prediction
}
