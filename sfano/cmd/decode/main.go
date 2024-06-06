package main

import (
	"compress/sfano/pkg/sfano"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <input file> <output file> <stats file>\n", os.Args[0])
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	statsFile := os.Args[3]

	sfano.DecodeFile(inputFile, outputFile, statsFile)
}
