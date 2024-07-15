package encode

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"slices"

	"github.com/gohuffman/frequency"
)

type childNode interface {
	Char() byte
	Freq() int
	IsLeaft() bool
}

type leaf struct {
	char byte
	freq int
}

func (l leaf) Char() byte {
	return l.char
}
func (l leaf) Freq() int {
	return l.freq
}
func (l leaf) IsLeaft() bool {
	return true
}

type baseNodeStruct struct {
	freq      int
	LeftNode  childNode
	RigthNode childNode
}

func (l baseNodeStruct) Char() byte {
	return 0
}
func (l baseNodeStruct) Freq() int {
	return l.freq
}
func (l baseNodeStruct) IsLeaft() bool {
	return false
}

type threeStruct struct {
	Root childNode
}

func (l threeStruct) Freq() int {
	return l.Root.Freq()
}

type plainT struct {
	Char byte
	Path byte
}

func (pt plainT) String() string {
	if pt.Char == 0 {
		return "0 "
	}
	return fmt.Sprintf("%c ", pt.Char)
}

type pathStruct struct {
	Byte byte
	Bits int
}

func NewThree(a, b childNode) threeStruct {
	return threeStruct{
		Root: &baseNodeStruct{
			LeftNode:  a,
			RigthNode: b,
			freq:      a.Freq() + b.Freq(),
		},
	}
}

func CreateThree(frequencies []frequency.FrequencyStruct) threeStruct {
	pq := make(PriorityQueue, len(frequencies))
	for i, v := range frequencies {
		pq[i] = threeStruct{
			Root: leaf{
				char: v.Char, freq: v.Frequency,
			},
		}
	}
	heap.Init(&pq)

	var a, b, result threeStruct

	for pq.Len() > 1 {
		a = heap.Pop(&pq).(threeStruct)
		b = heap.Pop(&pq).(threeStruct)
		result = NewThree(a.Root, b.Root)
		heap.Push(&pq, result)
	}

	return result
}

func traverse(node childNode, table map[byte]pathStruct, path byte, bits int) {
	if node.IsLeaft() {
		table[node.Char()] = pathStruct{Byte: path, Bits: bits}
		return
	}
	baseNode := node.(*baseNodeStruct)
	if baseNode.LeftNode != nil {
		traverse(baseNode.LeftNode, table, path<<1, bits+1)
	}

	if baseNode.RigthNode != nil {
		traverse(baseNode.RigthNode, table, path<<1+1, bits+1)
	}
}

func CreateTable(three threeStruct, size int) map[byte]pathStruct {
	table := make(map[byte]pathStruct, size)
	path := byte(0)

	root := three.Root

	traverse(root, table, path, 0)

	return table
}

func inorderTraversal(root childNode, table map[byte]pathStruct) []plainT {
	if root == nil {
		return []plainT{{}}
	}

	if root.IsLeaft() {
		return []plainT{
			{Char: root.Char(), Path: table[root.Char()].Byte},
		}
	}
	baseNode := root.(*baseNodeStruct)

	result := []plainT{}
	result = append(result, inorderTraversal(baseNode.LeftNode, table)...)
	result = append(result, plainT{})
	result = append(result, inorderTraversal(baseNode.RigthNode, table)...)

	return result
}

func WriteThree(three threeStruct, table map[byte]pathStruct, f *bufio.Writer) error {
	data := inorderTraversal(three.Root, table)

	slices.Reverse(data)
	for _, val := range data {
		_, err := f.WriteString(val.String())
		if err != nil {
			return err
		}
	}
	return f.Flush()
}

func WriteToFile(table map[byte]pathStruct, w *bufio.Writer, r *bufio.Reader) error {
	var path pathStruct
	var b byte
	var err error
	bits := 0
	for {
		char, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		path = table[char]

		for i := 0; i < path.Bits; i++ {
			b = (path.Byte << i) | b
			bits++
			if bits == 8 {
				err = w.WriteByte(b)
				if err != nil {
					return nil
				}

				b = 0
				bits = 0
			}
		}
	}

	if bits > 0 {
		err = w.WriteByte(b)
		if err != nil {
			return nil
		}
	}
	return w.Flush()
}
