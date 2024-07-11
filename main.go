package main

import (
	"bufio"
	"github.com/gohuffman/frequency"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(file)
	frequency.CalculateFrequency(buff)
}
