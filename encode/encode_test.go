package encode

import (
	"bufio"
	"os"
	"testing"

	"github.com/gohuffman/frequency"
)

func getTestFile(filename string) *bufio.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file)
}

func TestCreateThree(t *testing.T) {
	input := getTestFile("../frequency/freq.txt")

	// []frequency.FrequencyStruct{
	// 	{Char: 'c', Frequency: 1},
	// 	{Char: '\n', Frequency: 1},
	// 	{Char: 'b', Frequency: 2},
	// 	{Char: 'a', Frequency: 3},
	// }

	frequencies := frequency.CalculateFrequency(input)

	result := CreateThree(frequencies)

	if result.Freq() != 7 {
		t.Fatalf("wrong three frequency, expected=7 got=%d, Three %+v", result.Freq(), result)
	}

	left := result.Root.(*BaseNode).LeftNode.(Leaf)

	if left.char != 'a' || left.Freq() != 3 {
		t.Fatalf("wrong value expected=a got=%q freq expected=3 got=%d", left.char, left.freq)
	}

	rigth := result.Root.(*BaseNode).RigthNode.(*BaseNode)

	left = rigth.LeftNode.(Leaf)
	if left.char != 'b' || left.Freq() != 2 {
		t.Fatalf("wrong value expected=b got=%q freq expected=2 got=%d", left.char, left.freq)
	}
}

func TestCreateTable(t *testing.T) {
	input := getTestFile("../frequency/freq.txt")
	tests := []struct {
		Char byte
		Path []byte
	}{
		{Char: 'c', Path: []byte{1, 1, 0}},
		{Char: '\n', Path: []byte{1, 1, 1}},
		{Char: 'b', Path: []byte{1, 0}},
		{Char: 'a', Path: []byte{0}},
	}

	// []frequency.FrequencyStruct{
	// 	{Char: 'c', Frequency: 1},
	// 	{Char: '\n', Frequency: 1},
	// 	{Char: 'b', Frequency: 2},
	// 	{Char: 'a', Frequency: 3},
	// }

	frequencies := frequency.CalculateFrequency(input)

	three := CreateThree(frequencies)

	table := CreateTable(three, len(frequencies))

	for _, expected := range tests {
		path := table[expected.Char]
		isExpected := true
		for i, v := range expected.Path {
			if v != path[i] {
				isExpected = false
			}
		}
		if !isExpected {
			t.Logf("wrong path char %q want=%q got=%q", expected.Char, expected.Path, path)
		}
	}
}