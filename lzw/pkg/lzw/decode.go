package lzw

import (
	"bytes"
	"encoding/binary"
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

	decodedData := LZWDecode(data)
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

func LZWDecode(data []byte) []byte {
	dict := make(map[int16][]byte)
	for i := 0; i < 256; i++ {
		dict[int16(i)] = []byte{byte(i)}
	}

	var (
		currCode, prevCode int16
		currSeq            []byte
		output             []byte
		dictSize           int16 = 256
	)

	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &currCode)
	if err != nil {
		log.Fatalf("Failed to read initial code: %v\n", err)
	}
	output = append(output, dict[currCode]...)
	prevCode = currCode

	for reader.Len() > 0 {
		err = binary.Read(reader, binary.LittleEndian, &currCode)
		if err != nil {
			log.Fatalf("Failed to read code: %v\n", err)
		}
		if seq, ok := dict[currCode]; ok {
			currSeq = seq
		} else if currCode == dictSize {
			currSeq = append(dict[prevCode], dict[prevCode][0])
		} else {
			log.Fatalf("Invalid LZW code: %d\n", currCode)
		}
		output = append(output, currSeq...)
		if dictSize < MaxDictSize {
			dict[dictSize] = append(dict[prevCode], currSeq[0])
			dictSize++
		} else {
			// Reset dictionary
			dict = make(map[int16][]byte)
			for i := 0; i < 256; i++ {
				dict[int16(i)] = []byte{byte(i)}
			}
			dictSize = 256
		}
		prevCode = currCode
	}

	return output
}
