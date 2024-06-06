package huffman

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"time"
)

func DecodeFile(inputFile, outputFile, statsFileName string) {
	startTime := time.Now()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}

	var huffmanTree *HuffmanNode
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&huffmanTree); err != nil {
		log.Fatalf("Failed to decode Huffman tree: %v", err)
	}

	// Start decoding after the tree
	remainingData := buf.Bytes() // Data that remains after the tree data
	var decodedData bytes.Buffer
	node := huffmanTree
	totalBits := len(remainingData) * 8

	for i := 0; i < totalBits; i++ {
		byteIndex := i / 8
		bitIndex := 7 - (i % 8)
		if byteIndex >= len(remainingData) {
			break
		}
		if (remainingData[byteIndex]>>bitIndex)&1 == 1 {
			node = node.Right
		} else {
			node = node.Left
		}
		if node.Left == nil && node.Right == nil {
			decodedData.WriteByte(node.Character)
			node = huffmanTree
		}
		PrintProgress(i+1, totalBits, "Decoding")
	}

	err = ioutil.WriteFile(outputFile, decodedData.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v\n", err)
	}

	elapsedTime := time.Since(startTime)

	stats := Statistics{
		Operation:            DecodeOperation,
		InputFileName:        inputFile,
		OutputFileName:       outputFile,
		RunningTime:          elapsedTime.String(),
		InputChars:           len(remainingData),
		InputBytes:           len(remainingData),
		OutputChars:          decodedData.Len(),
		OutputBytes:          decodedData.Len(),
		DifferentCharCount:   len(BuildFrequencyTable(decodedData.Bytes())), // Optionally update this logic
		CharacterFrequencies: ConvertFrequenciesToReadable(BuildFrequencyTable(decodedData.Bytes())),
	}

	RecordStatistics(statsFileName, stats)
	log.Println("File decoded successfully.")
}
