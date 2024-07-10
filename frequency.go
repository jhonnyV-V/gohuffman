package main

import (
	"bufio"
	"io"
	"sort"
)

type FrequencyStruct struct {
	Char      byte
	Frequency int
}

func calculateFrequency(file *bufio.Reader) []FrequencyStruct {
	var frequencies []FrequencyStruct
	var freqmap map[byte]int = map[byte]int{}
	for {
		character, err := file.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		freqmap[character]++
	}

	for key, value := range freqmap {
		frequencies = append(frequencies, FrequencyStruct{Char: key, Frequency: value})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Frequency > frequencies[j].Frequency
	})
	return frequencies
}
