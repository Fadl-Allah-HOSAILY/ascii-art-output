package functions

import (
	"fmt"
	"os"
)

func WriteAsciiArt(asciiArtResult string, fileName string) {
	sliceOfByte := []byte(asciiArtResult)
	writeError := os.WriteFile(fileName, sliceOfByte, 0o644)
	if writeError != nil {
		fmt.Println("File write error!", writeError)
		return
	}
}
