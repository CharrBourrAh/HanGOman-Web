package asciimode

import (
	"log"
	"os"
	"strings"
)

//is not being used

func OpenAscii() [][]string {
	var AsciiTable [][]string
	AsciiFile, err := os.ReadFile("data/standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	AsciiStr := string(AsciiFile)
	AsciiSlice := strings.Split(AsciiStr, "\n")
	x := 1
	y := 0
	for x <= 27 {
		i := 0
		for x < x*8 {
			AsciiTable[i][x-1] = AsciiSlice[y]
			y++
			i++
		}
		x++

	}
	return AsciiTable
}

func ToAsciiArt(word string) string {
	asciiWord := ""
	AsciiTable := OpenAscii()
	var TableIndex []int
	for _, v := range word {
		if v == 95 {
			TableIndex = append(TableIndex, 0)
		} else {
			TableIndex = append(TableIndex, int(v-64))
		}
	}
	asciiWord += "\n"
	for i := 0; i < 8; i++ {
		for _, index := range TableIndex {
			asciiWord = asciiWord + "░░" + AsciiTable[i][index]
		}
		asciiWord += "░░"
		asciiWord += "\n"
	}
	return asciiWord
}
