package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"FinanransdonScript/FinanransdonScriptLib"
)

var args []string

func showHelp() {
	fmt.Printf("Vias Utendi: %s [<Optiones>] <Initus File> <Output File> \n", args[0])
	fmt.Printf(`Optiones:
	-h, --help, --auxilium, -a, -?: Praestare hoc auxilium nuntium.
	-v, --verbose, --verbosus: Plures informationes de executione monstrare.
	-d, --decompile: Decompile initus lima.

`)
}

func flagOn(flagNames ...string) bool {
	for _, i := range args {
		for _, p := range flagNames {
			if i == p {
				return true
			}
		}
	}
	return false
}

func main() {
	args = os.Args
	if len(args) < 3 {
		showHelp()
	} else {
		verbose := false
		if flagOn("-h", "--help", "--auxilium") {
			showHelp()
			return
		}
		if flagOn("-v", "--verbosus", "--verbose") {
			verbose = true
		}
		if flagOn("-d", "--decompile") {
			inputFilePath := args[len(args)-2]
			outputFilePath := args[len(args)-1]
			inputFile, err := os.Open(inputFilePath)
			if err != nil {
				fmt.Printf("Error dum foramen file: %s\n", inputFilePath)
			}
			bufferReader := bufio.NewReader(inputFile)
			content := []byte{}
			for {
				byteContent, err := bufferReader.ReadByte()
				if err != nil {
					break
				}
				content = append(content, byteContent)
			}
			inputFile.Close()
			result := FinanransdonScriptLib.DecompileFromData(content, verbose)
			err = ioutil.WriteFile(outputFilePath, []byte(result), 0644)
			if err != nil {
				fmt.Printf("Error dum foramen file: %s\n", outputFilePath)
			}
			return
		}
		inputFilePath := args[len(args)-2]
		outputFilePath := args[len(args)-1]
		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			fmt.Printf("Error dum foramen file: %s\n", inputFilePath)
		}
		bufferReader := bufio.NewReader(inputFile)
		content := ""
		for {
			byteContent, err := bufferReader.ReadByte()
			if err != nil {
				break
			}
			content += string(rune(byteContent))
		}
		inputFile.Close()
		byteArray := FinanransdonScriptLib.RunScript(content, verbose)
		if byteArray == nil {
			return
		}
		err = ioutil.WriteFile(outputFilePath, byteArray, 0644)
		if err != nil {
			fmt.Printf("Error dum foramen file: %s\n", outputFilePath)
		}
	}
}
