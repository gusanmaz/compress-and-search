package main

import (
	"compress/huffman/pkg/huffman"
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

	huffman.DecodeFile(inputFile, outputFile, statsFile)
}
