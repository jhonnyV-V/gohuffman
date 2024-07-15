package decode

import (
	"bufio"
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
