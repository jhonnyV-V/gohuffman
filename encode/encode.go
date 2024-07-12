package encode

import (
	"container/heap"
	"fmt"

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

func traverse(node ChildNode, table map[byte]byte, path byte) {
	if node.IsLeaft() {
		fmt.Printf("char %q, path %04b\n", node.Char(), path)
		table[node.Char()] = path
		return
	}
	baseNode := node.(*BaseNode)
	if baseNode.LeftNode != nil {
		fmt.Printf("go left prev path %0b\n", path)
		fmt.Printf("go left new path %0b\n", path<<1)
		traverse(baseNode.LeftNode, table, path<<1)
	}

	if baseNode.RigthNode != nil {
		fmt.Printf("go rigth prev path %0b\n", path)
		fmt.Printf("go rigth new path %0b\n", path<<1+1)
		traverse(baseNode.RigthNode, table, path<<1+1)
	}
}

func CreateTable(three Three, size int) map[byte]byte {
	table := make(map[byte]byte, size)
	path := byte(0)

	root := three.Root

	traverse(root, table, path)

	return table
}

// func WriteToFile(table map[byte]uint32, mapSize int, w *bufio.Writer, r *bufio.Reader) {
// 	size := uint32((mapSize * 5) + 4)
// 	headerBytes := make([]byte, size)
// 	//first bytes should be the size of the headers
// 	//it should be a version if I was mantaining this
// 	binary.LittleEndian.PutUint32(headerBytes, uint32(mapSize*5))
// 	fmt.Println("header size", len(headerBytes))
// 	fmt.Printf("header %#v\n", headerBytes)
//
// 	for char, path := range table {
// 		headerBytes = append(headerBytes, char)
// 		binary.LittleEndian.PutUint32(headerBytes[len(headerBytes)-1:], uint32(path))
// 	}
//
// 	w.Write(headerBytes)
// 	headerBytes = []byte{}
//
// 	for {
// 		char, err := r.ReadByte()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			} else {
// 				panic(err)
// 			}
// 		}
// 		//do the thing
// 		encoded := []byte{}
// 		binary.LittleEndian.PutUint32(encoded, table[char])
//
// 		w.Write(encoded)
// 	}
// }
