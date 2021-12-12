package main

import (
	"strings"
	"testing"
)

// indices into transliteration table
const ARABIC int = 0
const IPA int = 1

var transliterationTable = [][2]string{
	{"أ", "ʔ"},
	{"إ", "ʔ"},
	{"ء", "ʔ"},
	{"ؤ", "ʔ"},
	{"ئ", "ʔ"},
	{"ب", "b"},
	{"ت", "t"},
	{"ث", "θ"},
	{"ج", "d͜ʒ"},
	{"ح", "ħ"},
	{"خ", "χ"},
	{"د", "d"},
	{"ذ", "ð"},
	{"ر", "r"},
	{"ز", "z"},
	{"س", "s"},
	{"ش", "ʃ"},
	{"ص", "sˤ"},
	{"ض", "dˤ"},
	{"ط", "tˤ"},
	{"ظ", "ðˤ"},
	{"ع", "ʢ"},
	{"غ", "ʁ"},
	{"ف", "f"},
	{"ق", "q"},
	{"ك", "k"},
	{"ل", "l"},
	{"م", "m"},
	{"ن", "n"},
	{"ه", "h"},
	{"و", "w"},
	{"ي", "j"},
	{"ة", "t"},
	{"ى", "aː"},
}

// TransliterateArabicToIPA takes text written in
// the Arabic script and converts all consonants
// to International Phonetic Alphabet.
// Anything not recognized is not converted but
// copied into the resulting text.
func TransliterateArabicToIPA(arabicText string) string {
	var result string
	for len(arabicText) > 0 {
		var foundMatch bool

		for i := 0; i < len(transliterationTable); i++ {
			if strings.HasPrefix(arabicText, transliterationTable[i][ARABIC]) {
				result += transliterationTable[i][IPA]
				arabicText = strings.TrimPrefix(arabicText, transliterationTable[i][ARABIC])
				foundMatch = true
			}
		}

		if !foundMatch {
			var firstCharAsString string
			var arabicTextRunes []rune = []rune(arabicText)
			var firstChar rune = arabicTextRunes[0]
			firstCharAsString = string(firstChar)
			result += firstCharAsString
			arabicText = strings.TrimPrefix(arabicText, firstCharAsString)
		}

	}
	return result
}

func TestTransliterateArabicToIPA(t *testing.T) {
	tests := []struct {
		arabicText string
		want       string
	}{
		{arabicText: "شجرة", want: "ʃd͜ʒrt"},
		{arabicText: "أ", want: "ʔ"},
		{arabicText: "ب", want: "b"},
		{arabicText: "ت", want: "t"},
		{arabicText: "ثلج كثير", want: "θld͜ʒ kθjr"},
	}
	for _, tt := range tests {
		if got := TransliterateArabicToIPA(tt.arabicText); got != tt.want {
			t.Errorf("TransliterateArabicToIPA(%v) = %v, want %v", tt.arabicText, got, tt.want)
		}
	}
}

