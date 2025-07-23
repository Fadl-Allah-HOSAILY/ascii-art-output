package functions

import (
	"log"
	"os"
)

func WriteAsciiArt(asciiArtResult string, fileName string) {
	sliceOfByte := []byte(asciiArtResult)
	writeError := os.WriteFile(fileName, sliceOfByte, 0o644)
	if writeError != nil {
		log.Fatalln("Error", writeError)
	}
}