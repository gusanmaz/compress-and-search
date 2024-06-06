package sfano

import (
	"encoding/json"
	"log"
	"os"
)

type OperationType string

const (
	EncodeOperation OperationType = "encode"
	DecodeOperation OperationType = "decode"
)

type Statistics struct {
	Operation            OperationType  `json:"operation"`
	InputFileName        string         `json:"input_file_name"`
	OutputFileName       string         `json:"output_file_name"`
	RunningTime          string         `json:"running_time"`
	InputChars           int            `json:"input_chars"`
	InputBytes           int            `json:"input_bytes"`
	OutputChars          int            `json:"output_chars"`
	OutputBytes          int            `json:"output_bytes"`
	CompressionRate      float64        `json:"compression_rate,omitempty"`
	DifferentCharCount   int            `json:"different_char_count,omitempty"`
	CharacterFrequencies map[string]int `json:"character_frequencies,omitempty"`
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

func ConvertFrequenciesToReadable(frequencyTable map[byte]int) map[string]int {
	readableFreq := make(map[string]int)
	for k, v := range frequencyTable {
		readableFreq[string(rune(k))] = v
	}
	return readableFreq
}

func BuildFrequencyTable(data []byte) map[byte]int {
	frequencyTable := make(map[byte]int)
	for _, b := range data {
		frequencyTable[b]++
	}
	return frequencyTable
}
