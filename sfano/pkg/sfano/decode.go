package sfano

import (
	"bytes"
	"log"
	"time"
)

func DecodeFile(inputFile, outputFile, statsFileName string) {
	startTime := time.Now()
	data, err := readFileUTF8(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}
	log.Printf("Read %d bytes from input file\n", len(data))

	decodedData := ShannonFanoDecode(data)
	log.Printf("Decoded data size: %d bytes\n", len(decodedData))

	err = writeFileUTF8(outputFile, decodedData)
	if err != nil {
		log.Fatalf("Failed to write output file: %v\n", err)
	}

	elapsedTime := time.Since(startTime)

	stats := Statistics{
		Operation:            DecodeOperation,
		InputFileName:        inputFile,
		OutputFileName:       outputFile,
		RunningTime:          elapsedTime.String(),
		InputChars:           len(data),
		InputBytes:           len(data),
		OutputChars:          len(decodedData),
		OutputBytes:          len(decodedData),
		DifferentCharCount:   len(BuildFrequencyTable(decodedData)),
		CharacterFrequencies: ConvertFrequenciesToReadable(BuildFrequencyTable(decodedData)),
	}

	RecordStatistics(statsFileName, stats)
	log.Println("File decoded successfully.")
}

func ShannonFanoDecode(data []byte) []byte {
	// Rebuild frequency table and code table for decoding
	freqTable := BuildFrequencyTable(data)
	codeTable := BuildShannonFanoCodes(freqTable)

	var (
		currCode string
		output   bytes.Buffer
	)

	for _, b := range data {
		currCode += string(b)
		for k, v := range codeTable {
			if v == currCode {
				output.WriteByte(k)
				currCode = ""
				break
			}
		}
	}

	return output.Bytes()
}
