package frequency

import (
	"bufio"
	"os"
	"testing"
)

func getTestFile(filename string) *bufio.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file)
}

func TestGetFrequency(t *testing.T) {
	input := getTestFile("freq.txt")

	tests := []FrequencyStruct{
		{Char: 'c', Frequency: 1},
		{Char: '\n', Frequency: 1},
		{Char: 'b', Frequency: 2},
		{Char: 'a', Frequency: 3},
	}

	result := CalculateFrequency(input)

	for i, expected := range tests {
		actual := result[i]
		if actual.Char != expected.Char {
			t.Fatalf("got wrong char expected %q got %q\n", expected.Char, actual.Char)
		}

		if actual.Frequency != expected.Frequency {
			t.Fatalf("got wrong frequency for char %q: expected %d got %d\n", expected.Char, expected.Frequency, actual.Frequency)
		}
	}

}

func TestWithBigFile(t *testing.T) {
	input := getTestFile("../135-0.txt")

	result := CalculateFrequency(input)

	for _, actual := range result {
		if actual.Char == 't' {
			if actual.Frequency != 223000 {
				t.Fatalf("got wrong frequency for char t: expected 223000 got %d\n", actual.Frequency)
			}
		}

		if actual.Char == 'X' {
			if actual.Frequency != 333 {
				t.Fatalf("got wrong frequency for char X: expected 333 got %d\n", actual.Frequency)
			}
		}
	}
}
