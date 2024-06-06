package kmp

import (
	"log"
	"time"
)

func SearchFile(pattern, inputFile, statsFileName string) {
	startTime := time.Now()
	data, err := readFileUTF8(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}
	log.Printf("Read %d bytes from input file\n", len(data))

	matches := KMPSearch(pattern, string(data))
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

func KMPSearch(pattern, text string) []int {
	m := len(pattern)
	n := len(text)
	lps := computeLPSArray(pattern, m)
	var matches []int

	i, j := 0, 0
	for i < n {
		if pattern[j] == text[i] {
			j++
			i++
		}
		if j == m {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return matches
}

func computeLPSArray(pattern string, m int) []int {
	lps := make([]int, m)
	length := 0
	lps[0] = 0

	i := 1
	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}
