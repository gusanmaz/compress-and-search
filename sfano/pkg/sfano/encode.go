package sfano

import (
	"bytes"
	"log"
	"sort"
	"time"
)

func EncodeFile(inputFile, outputFile, statsFileName string) {
	startTime := time.Now()
	data, err := readFileUTF8(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}
	log.Printf("Read %d bytes from input file\n", len(data))

	encodedData := ShannonFanoEncode(data)
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

func ShannonFanoEncode(data []byte) []byte {
	freqTable := BuildFrequencyTable(data)
	codeTable := BuildShannonFanoCodes(freqTable)

	var output bytes.Buffer
	for _, b := range data {
		code := codeTable[b]
		for _, bit := range code {
			if bit == '0' {
				output.WriteByte(0)
			} else {
				output.WriteByte(1)
			}
		}
	}

	return output.Bytes()
}

func BuildShannonFanoCodes(freqTable map[byte]int) map[byte]string {
	type freqNode struct {
		symbol byte
		freq   int
		code   string
	}

	// Create a slice of freqNodes from the frequency table
	nodes := make([]freqNode, 0, len(freqTable))
	for k, v := range freqTable {
		nodes = append(nodes, freqNode{symbol: k, freq: v})
	}

	// Sort nodes by frequency in descending order
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].freq > nodes[j].freq
	})

	var buildCodes func(nodes []freqNode) []freqNode
	buildCodes = func(nodes []freqNode) []freqNode {
		if len(nodes) == 1 {
			return nodes
		}

		totalFreq := 0
		for _, node := range nodes {
			totalFreq += node.freq
		}

		sumFreq := 0
		var splitIndex int
		for i := 0; i < len(nodes); i++ {
			sumFreq += nodes[i].freq
			if sumFreq >= totalFreq/2 {
				splitIndex = i + 1
				break
			}
		}

		left := nodes[:splitIndex]
		right := nodes[splitIndex:]

		for i := range left {
			left[i].code += "0"
		}
		for i := range right {
			right[i].code += "1"
		}

		return append(buildCodes(left), buildCodes(right)...)
	}

	nodes = buildCodes(nodes)

	codeTable := make(map[byte]string)
	for _, node := range nodes {
		codeTable[node.symbol] = node.code
	}

	return codeTable
}
