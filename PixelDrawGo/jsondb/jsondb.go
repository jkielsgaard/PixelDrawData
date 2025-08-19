package jsondb

import (
	"bufio"
	"encoding/json"
	"os"
	"pixel-draw-go/pixeldata"
)

const dbFileName = "PixelTrainingData.jsonl"

// PutData marshals the given PixelJSON object and appends it to the database file.
func PutData(data pixeldata.PixelJSON) error {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(dbFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the JSON data followed by a newline
	if _, err := file.Write(append(jsonData, '\n')); err != nil {
		return err
	}

	return nil
}

// GetData reads all PixelJSON objects from the database file.
func GetData() ([]pixeldata.PixelJSON, error) {
	var pixelDataList []pixeldata.PixelJSON

	// Open the file for reading
	file, err := os.Open(dbFileName)
	if err != nil {
		// If the file doesn't exist, return an empty list, which is not an error condition
		if os.IsNotExist(err) {
			return pixelDataList, nil
		}
		return nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var pd pixeldata.PixelJSON
		// Unmarshal each line into a PixelJSON struct
		if err := json.Unmarshal(scanner.Bytes(), &pd); err != nil {
			// You might want to handle corrupted lines differently, e.g., log the error
			continue
		}
		pixelDataList = append(pixelDataList, pd)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pixelDataList, nil
}
