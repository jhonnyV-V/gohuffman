package main

import (
	"bufio"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	buff := bufio.NewReader(file)
	calculateFrequency(buff)
}
