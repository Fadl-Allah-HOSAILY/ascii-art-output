package functions

import (
	"log"
	"os"
)

func CheckLenght() {
	standard, errStandard := os.ReadFile("banners/standard.txt")
	if errStandard != nil {
		log.Fatalln("Error: ", errStandard)
	}

	shadow, errShadow := os.ReadFile("banners/shadow.txt")
	if errShadow != nil {
		log.Fatalln("Error: ", errShadow)
	}

	thinkertoy, errThinkertoy := os.ReadFile("banners/thinkertoy.txt")
	if errThinkertoy != nil {
		log.Fatalln("Error: ", errThinkertoy)
	}

	if len(standard) != 6622 {
		log.Fatalln("Error: banner standard is not correctly formated")
	}

	if len(shadow) != 7462 {
		log.Fatalln("Error: banner shadow is not correctly formated")
	}

	if len(thinkertoy) != 4702 {
		log.Fatalln("Error: banner thinkertoy is not correctly formated")
	}
}
