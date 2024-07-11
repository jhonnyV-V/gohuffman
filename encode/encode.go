package encode

import (
	"container/heap"

	"github.com/gohuffman/frequency"
)

type ChildNode interface {
	Char() byte
	Freq() int
	IsLeaft() bool
}

type Leaf struct {
	char byte
	freq int
}

func (l Leaf) Char() byte {
	return l.char
}
func (l Leaf) Freq() int {
	return l.freq
}
func (l Leaf) IsLeaft() bool {
	return true
}

type BaseNode struct {
	freq      int
	LeftNode  ChildNode
	RigthNode ChildNode
}

func (l BaseNode) Char() byte {
	return 0
}
func (l BaseNode) Freq() int {
	return l.freq
}
func (l BaseNode) IsLeaft() bool {
	return false
}

type Three struct {
	Root ChildNode
}

func (l Three) Freq() int {
	return l.Root.Freq()
}

func NewThree(a, b ChildNode) Three {
	return Three{
		Root: &BaseNode{
			LeftNode:  a,
			RigthNode: b,
			freq:      a.Freq() + b.Freq(),
		},
	}
}

func CreateThree(frequencies []frequency.FrequencyStruct) Three {
	pq := make(PriorityQueue, len(frequencies))
	for i, v := range frequencies {
		pq[i] = Three{
			Root: Leaf{
				char: v.Char, freq: v.Frequency,
			},
		}
	}
	heap.Init(&pq)

	var a, b, result Three

	for pq.Len() > 1 {
		a = heap.Pop(&pq).(Three)
		b = heap.Pop(&pq).(Three)
		result = NewThree(a.Root, b.Root)
		heap.Push(&pq, result)
	}

	return result
}

func traverse(node ChildNode, table map[byte][]byte, path []byte) {
	if node.IsLeaft() {
		table[node.Char()] = path
		return
	}
	baseNode := node.(*BaseNode)
	if baseNode.LeftNode != nil {
		newPath := append(path, 0)
		traverse(baseNode.LeftNode, table, newPath)
	}

	if baseNode.RigthNode != nil {
		newPath := make([]byte, len(path))
		copy(newPath, path)
		newPath = append(newPath, 1)
		traverse(baseNode.RigthNode, table, newPath)
	}
}

func CreateTable(three Three, size int) map[byte][]byte {
	table := make(map[byte][]byte, size)
	path := []byte{}

	root := three.Root

	traverse(root, table, path)

	return table
}
