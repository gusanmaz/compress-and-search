package karprabin

import (
	"bufio"
	"fmt"
	"os"
)

func readFileUTF8(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var output []byte
	buffer := make([]byte, 4096)
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			break
		}
		output = append(output, buffer[:bytesRead]...)
	}
	return output, nil
}

func PrintProgress(current, total int, process string) {
	percent := (current * 100) / total
	fmt.Printf("\r%s: %d%% complete", process, percent)
	if current == total {
		fmt.Println()
	}
}
