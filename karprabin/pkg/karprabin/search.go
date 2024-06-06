package karprabin

import (
	"hash/fnv"
	"log"
	"time"
)

const primeRK = 16777619

func SearchFile(pattern, inputFile, statsFileName string) {
	startTime := time.Now()
	data, err := readFileUTF8(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}
	log.Printf("Read %d bytes from input file\n", len(data))

	matches := KarpRabinSearch(pattern, string(data))
	log.Printf("Found %d matches at positions: %v\n", len(matches), matches)

	elapsedTime := time.Since(startTime)

	stats := Statistics{
		Operation:     SearchOperation,
		InputFileName: inputFile,
		Pattern:       pattern,
		RunningTime:   elapsedTime.String(),
		Matches:       matches,
	}

	RecordStatistics(statsFileName, stats)
	log.Println("Search completed successfully.")
}

func KarpRabinSearch(pattern, text string) []int {
	m := len(pattern)
	hpattern := hashPattern(pattern)
	h := fnv.New32()
	var matches []int

	for i := 0; i < len(text)-m+1; i++ {
		h.Write([]byte(text[i : i+m]))
		if h.Sum32() == hpattern && text[i:i+m] == pattern {
			matches = append(matches, i)
		}
		h.Reset()
	}

	return matches
}

func hashPattern(pattern string) uint32 {
	h := fnv.New32()
	h.Write([]byte(pattern))
	return h.Sum32()
}
