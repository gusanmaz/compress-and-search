package huffman

import "fmt"

func PrintProgress(current, total int, process string) {
	percent := (current * 100) / total
	fmt.Printf("\r%s: %d%% complete", process, percent)
	if current == total {
		fmt.Println()
	}
}
