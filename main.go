package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"youtube.com.fake/GrassDewMorning/arabicToIPA/transliteration"
)

// This function reads all lines from
// standard input, transliterating each
// one at a time.
func standardInToStandardOut(transliterate func(string) string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var currentLine string = scanner.Text()
		var transliteratedLine string = transliterate(currentLine)
		fmt.Println(transliteratedLine)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

func main() {
	standardInToStandardOut(transliteration.TransliterateArabicToIPA)
}
