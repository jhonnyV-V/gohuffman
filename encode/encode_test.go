package encode

import (
	"bufio"
	"os"
	"strings"
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
		Path byte
	}{
		{Char: 'c', Path: 0b110},
		{Char: '\n', Path: 0b111},
		{Char: 'b', Path: 0b10},
		{Char: 'a', Path: 0b0},
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
		if path != expected.Path {
			t.Fatalf("wrong path char=%q want=%#v got=%#v", expected.Char, expected.Path, path)
		}
	}
}

func TestWriteThree(t *testing.T) {
	input := getTestFile("../frequency/freq.txt")

	// []frequency.FrequencyStruct{
	// 	{Char: 'c', Frequency: 1},
	// 	{Char: '\n', Frequency: 1},
	// 	{Char: 'b', Frequency: 2},
	// 	{Char: 'a', Frequency: 3},
	// }

	frequencies := frequency.CalculateFrequency(input)

	three := CreateThree(frequencies)

	table := CreateTable(three, len(frequencies))

	fthree := new(strings.Builder)

	threeBuff := bufio.NewWriter(fthree)

	err := WriteThree(three, table, threeBuff)
	if err != nil {
		panic(err)
	}

	expected := []string{"'\\n':111", "0", "'c':110", "0", "'b':10", "0", "'a':0"}
	actual := strings.Split(strings.TrimSpace(fthree.String()), " ")

	if len(actual) != len(expected) {
		t.Fatalf("wrong three len want=%d, got=%d\ngot=%+v\nwant=%+v\n", len(expected), len(actual), actual, expected)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("wrong value want=%s got=%s", expected[i], actual[i])
		}
	}
}
