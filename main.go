package main

import (
	"fmt"

	"asciiArtOutput/functions"
)

func main() {
	fileName, inputText, banner, checkOutput, checkError := functions.ReadArgs()
	if checkError {
		return
	}

	if !functions.CheckRange(inputText) {
		fmt.Println("error, string is not valid ")
		return
	}

	if !functions.CheckBanner(banner) {
		fmt.Println("error, banner is not valid ")
		return
	}

	// split input text to slice of string
	wordsSlice := functions.SplitInput(inputText)

	// calling functions to handle the input
	functions.AppendArt(fileName, wordsSlice, functions.AsciiArtTable(banner), checkOutput)
}
