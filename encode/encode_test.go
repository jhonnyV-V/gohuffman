package encode

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/gohuffman/frequency"
)

func getTestFile(filename string) (*bufio.Reader, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file), file
}

func TestCreateThree(t *testing.T) {
	input, _ := getTestFile("../frequency/freq.txt")
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

	left := result.Root.(*baseNodeStruct).LeftNode.(leaf)

	if left.char != 'a' || left.Freq() != 3 {
		t.Fatalf("wrong value expected=a got=%q freq expected=3 got=%d", left.char, left.freq)
	}

	rigth := result.Root.(*baseNodeStruct).RigthNode.(*baseNodeStruct)

	left = rigth.LeftNode.(leaf)
	if left.char != 'b' || left.Freq() != 2 {
		t.Fatalf("wrong value expected=b got=%q freq expected=2 got=%d", left.char, left.freq)
	}
}

func TestCreateTable(t *testing.T) {
	input, _ := getTestFile("../frequency/freq.txt")
	tests := []struct {
		Char byte
		Path byte
		Bits int
	}{
		{Char: 'c', Path: 0b110, Bits: 3},
		{Char: '\n', Path: 0b111, Bits: 3},
		{Char: 'b', Path: 0b10, Bits: 2},
		{Char: 'a', Path: 0b0, Bits: 1},
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
		if path.Byte != expected.Path {
			t.Fatalf("wrong path char=%q want=%#v got=%#v", expected.Char, expected.Path, path)
		}
		if path.Bits != expected.Bits {
			t.Fatalf("wrong bits char=%q want=%d got=%d", expected.Char, expected.Bits, path.Bits)
		}
	}
}

func TestWriteThree(t *testing.T) {
	input, _ := getTestFile("../frequency/freq.txt")
	frequencies := frequency.CalculateFrequency(input)
	three := CreateThree(frequencies)
	table := CreateTable(three, len(frequencies))

	fthree := new(strings.Builder)
	threeBuff := bufio.NewWriter(fthree)

	err := WriteThree(three, table, threeBuff)
	if err != nil {
		panic(err)
	}

	expected := []string{"\n", "0", "c", "0", "b", "0", "a"}
	actual := strings.Split(strings.TrimSuffix(fthree.String(), " "), " ")

	if len(actual) != len(expected) {
		t.Fatalf("wrong three len want=%d, got=%d\ngot=%+v\nwant=%+v\n", len(expected), len(actual), actual, expected)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("wrong value want=%s got=%s", expected[i], actual[i])
		}
	}
}

func TestWriteEncoded(t *testing.T) {
	input, _ := getTestFile("../frequency/freq.txt")

	// aaabbc\n
	frequencies := frequency.CalculateFrequency(input)
	three := CreateThree(frequencies)
	table := CreateTable(three, len(frequencies))

	outFile := new(bytes.Buffer)
	outBuff := bufio.NewWriter(outFile)

	input2, _ := getTestFile("../frequency/freq.txt")

	err := WriteToFile(table, outBuff, input2)
	if err != nil {
		panic(err)
	}

	//00010101 10111000
	expectedBytes := []byte{0b00010101, 0b10111000}
	actual := outFile.Bytes()

	t.Logf("%#v", actual)

	if len(actual) != len(expectedBytes) {
		t.Fatalf(
			"wrong length want=%d got=%d expected=%#v actual=%#v",
			len(expectedBytes),
			len(actual),
			expectedBytes,
			actual,
		)
	}

	for i, expected := range expectedBytes {
		if expected == actual[i] {
			t.Fatalf("wrong value want=%0b got=%0b", expected, actual)
		}
	}
}
