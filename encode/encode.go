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

func traverse(node ChildNode, table map[byte]int32, path int32) int32 {
	if node.IsLeaft() {
		table[node.Char()] = path
		return path + int32(1)
	}
	baseNode := node.(*BaseNode)
	localPath := path
	if baseNode.LeftNode != nil {
		localPath = traverse(baseNode.LeftNode, table, path)
	}

	if baseNode.RigthNode != nil {
		localPath = traverse(baseNode.RigthNode, table, localPath)
	}

	return localPath
}

func CreateTable(three Three, size int) map[byte]int32 {
	table := make(map[byte]int32, size)
	path := int32(0)

	root := three.Root

	traverse(root, table, path)

	return table
}
