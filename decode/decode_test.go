package decode

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

func getTestFile(filename string) (*bufio.Reader, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file), file
}

func TestCreateThree(t *testing.T) {
	rawInput := "c 0 \n 0 b 0 a "
	buff := bufio.NewReader(strings.NewReader(rawInput))

	result := BuildThree(buff)

	t.Logf("%#v", result)

	if result.LeftNode.Char() != 'a' {
		t.Fatalf("wrong three, expected=a got=%c, Three %+v", result.LeftNode.Char(), result)
	}

	result = result.RigthNode.(*baseNodeStruct)

	if result.LeftNode.Char() != 'b' {
		t.Fatalf("wrong three, expected=b got=%c, Three %+v", result.LeftNode.Char(), result)
	}

	result = result.RigthNode.(*baseNodeStruct)

	if result.LeftNode.Char() != 'c' {
		t.Fatalf("wrong three, expected=c got=%c, Three %+v", result.LeftNode.Char(), result)
	}

	if result.RigthNode.Char() != '\n' {
		t.Fatalf("wrong three, expected=\\n got=%c, Three %+v", result.LeftNode.Char(), result)
	}
}

func TestDecodeFile(t *testing.T) {
	rawThree := "c 0 \n 0 b 0 a "
	buffThree := bufio.NewReader(strings.NewReader(rawThree))
	three := BuildThree(buffThree)

	rawBytes := []byte{0b00010101, 0b10111000}
	file := bytes.NewReader(rawBytes)
	r := bufio.NewReader(file)
	rawoutput := []byte{}
	outBuffer := bytes.NewBuffer(rawoutput)
	w := bufio.NewWriter(outBuffer)

	DecodeFile(three, r, w)

	tests := []byte("aaabbc\n")
	rawoutput = outBuffer.Bytes()

	for i, expected := range tests {
		if expected != rawoutput[i] {
			t.Fatalf("wrong char want=%c got=%c", expected, rawoutput[i])
		}
	}
}
