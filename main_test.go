package main

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
		{Char: 'a', Frequency: 3},
		{Char: 'b', Frequency: 2},
		{Char: 'c', Frequency: 1},
	}

	result := calculateFrequency(input)

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
