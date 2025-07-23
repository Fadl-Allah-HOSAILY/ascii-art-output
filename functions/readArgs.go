package functions

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadArgs() (string, string, string, bool, bool) {
	args := os.Args[1:]
	inputText := ""
	banner := ""
	fileName := ""
	checkOutput := true
	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return fileName, inputText, banner, checkOutput, true
	}
	switch len(args) {
	case 1:
		inputText = args[0]
		banner = "standard"
		checkOutput = false
	case 2:
		firstArg := args[0]
		secondtArg := args[1]
		if strings.HasPrefix(firstArg, "--output=") {
			if strings.Count(firstArg, "=") != 1 {
				log.Fatalln("Error: this is not the correct way to write a folder.")
			}
			fileName = strings.TrimLeft(firstArg, "--output=")
			if strings.Contains(fileName, "banners/standard.txt") || strings.Contains(fileName, "banners/shadow.txt") || strings.Contains(fileName, "banners/thinkertoy.txt") {
				log.Fatalln("Error: you cannot write to this file.")
			}
			if !strings.HasSuffix(fileName, ".txt") {
				log.Fatalln("Error: this is not text file!")
			}
			inputText = secondtArg
			banner = "standard"
			checkOutput = true
		} else {
			inputText = firstArg
			banner = secondtArg
			checkOutput = false
		}
	case 3:
		if strings.HasPrefix(args[0], "--output=") {
			if strings.Count(args[0], "=") != 1 {
				log.Fatalln("Error: this is not the correct way to write a folder.")
			}
			fileName = strings.TrimLeft(args[0], "--output=")
			if strings.Contains(fileName, "banners/standard.txt") || strings.Contains(fileName, "banners/shadow.txt") || strings.Contains(fileName, "banners/thinkertoy.txt") {
				log.Fatalln("Error: you cannot write to this file.")
			}
			if !strings.HasSuffix(fileName, ".txt") {
				log.Fatalln("Error: this is not text file!")
			}
			inputText = args[1]
			banner = args[2]
			checkOutput = true
		} else {
			log.Fatalln("Error: this is not the correct way to write a folder.")
		}
	}
	return fileName, inputText, banner, checkOutput, false
}