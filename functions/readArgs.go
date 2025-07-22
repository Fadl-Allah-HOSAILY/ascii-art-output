package functions

import (
	"fmt"
	"os"
	"strings"
)

func ReadArgs() (string, string, string, bool, bool) {
	// checking the validity of input
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
			fileName = strings.TrimLeft(firstArg, "--output=")
			if strings.Contains(fileName, "banners/standard.txt") || strings.Contains(fileName, "banners/shadow.txt") || strings.Contains(fileName, "banners/thinkertoy.txt") {
				fmt.Println("error: -----------------!")
				return fileName, inputText, banner, checkOutput, true
			}
			if !strings.HasSuffix(fileName, ".txt") {
				fmt.Println("error: the file is not text file!")
				return fileName, inputText, banner, checkOutput, true
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
			fileName = strings.TrimLeft(args[0], "--output=")
			if strings.Contains(fileName, "banners/standard.txt") || strings.Contains(fileName, "banners/shadow.txt") || strings.Contains(fileName, "banners/thinkertoy.txt") {
				fmt.Println("error: -----------------!")
				return fileName, inputText, banner, checkOutput, true
			}
			if !strings.HasSuffix(fileName, ".txt") {
				fmt.Println("error: the file is not text file!")
				return fileName, inputText, banner, checkOutput, true
			}
			inputText = args[1]
			banner = args[2]
			checkOutput = true
		} else {
			fmt.Println("error: the file is not text file!")
			return fileName, inputText, banner, checkOutput, true
		}
	}
	return fileName, inputText, banner, checkOutput, false
}
