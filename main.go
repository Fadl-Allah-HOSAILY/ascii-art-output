package main

import (
	"log"

	"asciiArtOutput/functions"
)

func main() {
	functions.CheckLenght()
	fileName, inputText, banner, checkOutput, checkError := functions.ReadArgs()
	if checkError {
		return
	}

	if !functions.CheckRange(inputText) {
		log.Fatalln("Error, string is not valid")
	}

	if !functions.CheckBanner(banner) {
		log.Fatalln("Error, banner is not valid ")
	}

	// split input text to slice of string
	wordsSlice := functions.SplitInput(inputText)

	// calling functions to handle the input
	functions.AppendArt(fileName, wordsSlice, functions.AsciiArtTable(banner), checkOutput)
}
