package main

import (
	"karprabin/pkg/karprabin"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <pattern> <input file> <stats file>\n", os.Args[0])
	}
	pattern := os.Args[1]
	inputFile := os.Args[2]
	statsFile := os.Args[3]

	karprabin.SearchFile(pattern, inputFile, statsFile)
}
