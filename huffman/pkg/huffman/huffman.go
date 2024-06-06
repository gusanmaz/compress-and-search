package huffman

import (
	"container/heap"
)

type HuffmanNode struct {
	Character byte
	Frequency int
	Left      *HuffmanNode
	Right     *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Frequency < pq[j].Frequency }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*HuffmanNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func BuildFrequencyTable(data []byte) map[byte]int {
	frequencyTable := make(map[byte]int)
	for _, b := range data {
		frequencyTable[b]++
	}
	return frequencyTable
}

func BuildHuffmanTree(frequencyTable map[byte]int) *HuffmanNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for char, freq := range frequencyTable {
		heap.Push(&pq, &HuffmanNode{Character: char, Frequency: freq})
	}
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		merged := &HuffmanNode{Frequency: left.Frequency + right.Frequency, Left: left, Right: right}
		heap.Push(&pq, merged)
	}
	return heap.Pop(&pq).(*HuffmanNode)
}

func BuildHuffmanCodes(node *HuffmanNode, code string, huffmanCodes map[byte]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		huffmanCodes[node.Character] = code
		return
	}
	BuildHuffmanCodes(node.Left, code+"0", huffmanCodes)
	BuildHuffmanCodes(node.Right, code+"1", huffmanCodes)
}

func EncodeData(data []byte, huffmanCodes map[byte]string) string {
	var encodedData string
	for _, b := range data {
		encodedData += huffmanCodes[b]
	}
	return encodedData
}

func DecodeData(encodedData string, root *HuffmanNode) []byte {
	var decodedData []byte
	node := root
	for _, bit := range encodedData {
		if bit == '0' {
			node = node.Left
		} else {
			node = node.Right
		}
		if node.Left == nil && node.Right == nil {
			decodedData = append(decodedData, node.Character)
			node = root
		}
	}
	return decodedData
}
