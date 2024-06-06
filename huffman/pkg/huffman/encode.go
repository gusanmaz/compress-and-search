package huffman

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"time"
)

func EncodeFile(inputFile, outputFile, statsFileName string) {
	startTime := time.Now()
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}

	frequencyTable := BuildFrequencyTable(data)
	huffmanTree := BuildHuffmanTree(frequencyTable)

	huffmanCodes := make(map[byte]string)
	BuildHuffmanCodes(huffmanTree, "", huffmanCodes)

	var encodedData bytes.Buffer
	encodeTree(&encodedData, huffmanTree) // Encode the Huffman tree first

	totalBytes := len(data)
	var currentByte byte
	var bitCount int

	for i, b := range data {
		code := huffmanCodes[b]
		for _, bit := range code {
			if bit == '1' {
				currentByte |= 1 << (7 - bitCount)
			}
			bitCount++
			if bitCount == 8 {
				encodedData.WriteByte(currentByte)
				currentByte = 0
				bitCount = 0
			}
		}
		PrintProgress(i+1, totalBytes, "Encoding")
	}
	if bitCount > 0 {
		encodedData.WriteByte(currentByte)
	}

	err = ioutil.WriteFile(outputFile, encodedData.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v\n", err)
	}

	elapsedTime := time.Since(startTime)
	compressionRate := float64(encodedData.Len()) / float64(len(data))

	stats := Statistics{
		Operation:            EncodeOperation,
		InputFileName:        inputFile,
		OutputFileName:       outputFile,
		RunningTime:          elapsedTime.String(),
		InputChars:           len(data),
		InputBytes:           len(data),
		OutputChars:          encodedData.Len(),
		OutputBytes:          encodedData.Len(),
		CompressionRate:      compressionRate,
		DifferentCharCount:   len(frequencyTable),
		CharacterFrequencies: ConvertFrequenciesToReadable(frequencyTable),
	}

	RecordStatistics(statsFileName, stats)
	log.Println("File encoded successfully.")
}

// Serialize the Huffman Tree using Gob
func encodeTree(buf *bytes.Buffer, root *HuffmanNode) {
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(root); err != nil {
		log.Fatalf("Failed to encode Huffman tree: %v", err)
	}
}
