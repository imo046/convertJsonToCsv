package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func run(input string, output string) {
	// Open the JSON file
	jsonFile, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	// Decode the JSON data

	var jsonData map[string]interface{}
	if err := json.Unmarshal(jsonFile, &jsonData); err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}

	// Define column names
	columnNames := []string{"ID", "Text_Session1"}

	// Create a CSV file
	csvFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV writer
	csvWriter := csv.NewWriter(csvFile)

	// Write header row
	if err := csvWriter.Write(columnNames); err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Write data rows
	for key, value := range jsonData {
		// Create a new row with key-value pair
		csvRow := []string{key, fmt.Sprintf("%v", value)}
		if err := csvWriter.Write(csvRow); err != nil {
			fmt.Println("Error writing CSV row:", err)
			return
		}

	}

	csvWriter.Flush()

	fmt.Println("Conversion completed. CSV data written to output.csv")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input output")
		os.Exit(1)
	}
	filePath := os.Args[1]
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {

		}
	}
	outPath := os.Args[2]

	if len(outPath) <= 0 {
		log.Fatal("Output cannot be empty")
	}

	run(filePath, outPath)

}
