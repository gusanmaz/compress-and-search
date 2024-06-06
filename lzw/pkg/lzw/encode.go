package lzw

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

func EncodeFile(inputFile, outputFile, statsFileName string) {
	startTime := time.Now()
	data, err := readFileUTF8(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}
	log.Printf("Read %d bytes from input file\n", len(data))

	encodedData := LZWEncode(data)
	log.Printf("Encoded data size: %d bytes\n", len(encodedData))

	err = writeFileUTF8(outputFile, encodedData)
	if err != nil {
		log.Fatalf("Failed to write output file: %v\n", err)
	}

	elapsedTime := time.Since(startTime)
	compressionRate := float64(len(encodedData)) / float64(len(data))

	stats := Statistics{
		Operation:            EncodeOperation,
		InputFileName:        inputFile,
		OutputFileName:       outputFile,
		RunningTime:          elapsedTime.String(),
		InputChars:           len(data),
		InputBytes:           len(data),
		OutputChars:          len(encodedData),
		OutputBytes:          len(encodedData),
		CompressionRate:      compressionRate,
		DifferentCharCount:   len(BuildFrequencyTable(data)),
		CharacterFrequencies: ConvertFrequenciesToReadable(BuildFrequencyTable(data)),
	}

	RecordStatistics(statsFileName, stats)
	log.Println("File encoded successfully.")
}

func LZWEncode(data []byte) []byte {
	dict := make(map[string]int16)
	for i := 0; i < 256; i++ {
		dict[string(byte(i))] = int16(i)
	}

	var (
		currSeq  string
		output   bytes.Buffer
		dictSize int16 = 256
	)

	for _, b := range data {
		currSeqPlusChar := currSeq + string(b)
		if _, ok := dict[currSeqPlusChar]; ok {
			currSeq = currSeqPlusChar
		} else {
			err := binary.Write(&output, binary.LittleEndian, dict[currSeq])
			if err != nil {
				log.Fatalf("Failed to write binary data: %v\n", err)
			}
			if dictSize < MaxDictSize {
				dict[currSeqPlusChar] = dictSize
				dictSize++
			} else {
				// Reset dictionary
				dict = make(map[string]int16)
				for i := 0; i < 256; i++ {
					dict[string(byte(i))] = int16(i)
				}
				dictSize = 256
			}
			currSeq = string(b)
		}
	}
	if currSeq != "" {
		err := binary.Write(&output, binary.LittleEndian, dict[currSeq])
		if err != nil {
			log.Fatalf("Failed to write binary data: %v\n", err)
		}
	}

	return output.Bytes()
}
