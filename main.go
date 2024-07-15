package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gohuffman/decode"
	"github.com/gohuffman/encode"
	"github.com/gohuffman/frequency"
)

func addPrefixToFilename(path string, prefix string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	newBase := prefix + base
	return filepath.Join(dir, newBase)
}

func main() {
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

	outputfilename := flag.Arg(2)
	if outputfilename == "" {
		fmt.Println("No Output File")
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

		threeFile, err := os.Create(addPrefixToFilename(outputfilename, "three-"))
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

	case "decode":
		threeFile, err := os.Open(addPrefixToFilename(filename, "three-"))
		if err != nil {
			panic(err)
		}
		threeBuff := bufio.NewReader(threeFile)
		three := decode.BuildThree(threeBuff)

		file, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
		buff = bufio.NewReader(file)

		decode.DecodeFile(three, buff, outBuff)
	}
}
