package decode

import (
	"bufio"
	"io"
)

type childNode interface {
	Char() byte
	IsLeaft() bool
}

type leaf struct {
	char byte
}

func (l leaf) Char() byte {
	return l.char
}
func (l leaf) IsLeaft() bool {
	return true
}

type baseNodeStruct struct {
	LeftNode  childNode
	RigthNode childNode
}

func (l baseNodeStruct) Char() byte {
	return 0
}
func (l baseNodeStruct) IsLeaft() bool {
	return false
}

func Traverse(node childNode, path byte) childNode {
	if node.IsLeaft() {
		return node
	}
	baseNode := node.(*baseNodeStruct)

	if path&0b1 == 0b1 {
		return baseNode.RigthNode
	}

	return baseNode.LeftNode
}

func readPair(buff *bufio.Reader) (byte, error) {
	var char byte

	char, err := buff.ReadByte()
	if err != nil {
		return 0, err
	}

	return char, nil
}

func readBaseNode(buff *bufio.Reader) (*baseNodeStruct, error) {
	_, err := buff.ReadByte()
	if err != nil {
		return nil, err
	}

	_, err = buff.ReadByte()
	if err != nil {
		return nil, err
	}

	_, err = buff.ReadByte()
	if err != nil {
		return nil, err
	}

	return &baseNodeStruct{}, nil
}

func BuildThree(buff *bufio.Reader) *baseNodeStruct {
	root := &baseNodeStruct{}
	var r childNode

	cha, err := readPair(buff)
	if err != nil {
		panic(err)
	}

	root, err = readBaseNode(buff)
	if err != nil {
		panic(err)
	}

	root.LeftNode = leaf{char: cha}

	cha, err = readPair(buff)
	if err != nil {
		panic(err)
	}
	root.RigthNode = leaf{char: cha}

	for {
		r = root
		root, err = readBaseNode(buff)
		if err != nil {
			if err == io.EOF {
				root = r.(*baseNodeStruct)
				break
			}
			panic(err)
		}
		root.RigthNode = r

		cha, err = readPair(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		root.LeftNode = &leaf{char: cha}
	}

	return root
}

func DecodeFile(three childNode, r *bufio.Reader, w *bufio.Writer) {
	var result []byte
	var lastNode childNode = three

	for {
		packedByte, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		for packedByte > 0 {
			if packedByte == 0 {
				continue
			}
			//get the leftmost bit
			path := packedByte >> 7

			// Clear the MSB
			packedByte &= 0b01111111
			// Shift left
			packedByte <<= 1
			lastNode = Traverse(lastNode, path)
			if lastNode.IsLeaft() {
				result = append(result, lastNode.Char())
				lastNode = three
			}
		}
	}

	_, err := w.Write(result)
	if err != nil {
		panic(result)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
