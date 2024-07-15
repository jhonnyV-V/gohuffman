package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/gohuffman/encode"
	"github.com/gohuffman/frequency"
)

func main() {
	outputfilename := ""
	flag.StringVar(&outputfilename, "output", "output.txt", "-output filename")
	flag.Parse()

	command := flag.Arg(0)
	if command != "encode" && command != "decode" {
		fmt.Println("command " + command)
		fmt.Println("Invalid Command, try with decode or encode")
		os.Exit(0)
	}

	filename := flag.Arg(1)
	if filename == "" {
		fmt.Println("No Filename")
		os.Exit(0)
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	outfile, err := os.Create(outputfilename)
	if err != nil {
		panic(err)
	}

	buff := bufio.NewReader(file)

	outBuff := bufio.NewWriter(outfile)

	switch command {
	case "encode":
		freq := frequency.CalculateFrequency(buff)
		three := encode.CreateThree(freq)
		table := encode.CreateTable(three, len(freq))
		buff.Reset(file)

		threeFile, err := os.Create("three-" + outputfilename)
		if err != nil {
			panic(err)
		}
		threeBuff := bufio.NewWriter(threeFile)

		err = encode.WriteThree(three, table, threeBuff)
		if err != nil {
			panic(err)
		}

		file, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
		buff = bufio.NewReader(file)

		err = encode.WriteToFile(table, outBuff, buff)
		if err != nil {
			panic(err)
		}
	}
}
