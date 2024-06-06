package karprabin

import (
	"encoding/json"
	"log"
	"os"
)

type OperationType string

const (
	SearchOperation OperationType = "search"
)

type Statistics struct {
	Operation     OperationType `json:"operation"`
	InputFileName string        `json:"input_file_name"`
	Pattern       string        `json:"pattern"`
	RunningTime   string        `json:"running_time"`
	Matches       []int         `json:"matches"`
}

type StatisticsCollection struct {
	Entries []Statistics `json:"entries"`
}

func RecordStatistics(statsFileName string, stats Statistics) {
	var statsCollection StatisticsCollection

	file, err := os.OpenFile(statsFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open stats file: %v", err)
	}
	defer file.Close()

	// Read existing stats
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&statsCollection); err != nil && err.Error() != "EOF" {
		log.Fatalf("Failed to decode stats: %v", err)
	}

	// Add new stats
	statsCollection.Entries = append(statsCollection.Entries, stats)

	// Write updated stats
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(statsCollection); err != nil {
		log.Fatalf("Failed to write stats: %v", err)
	}
}
