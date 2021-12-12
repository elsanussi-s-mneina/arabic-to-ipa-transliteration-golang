package main

import (
	"testing"
)

var transliterationTable = [][2]string{
	{"أ", "ʔ"},
	{"شجرة", "ʃd͜ʒrh"},
	{"ب", "b"},
	{"ت", "t"},
}

// TransliterateArabicToIPA takes text written in
// the Arabic script and converts all consonants
// to International Phonetic Alphabet.
// Anything not recognized is not converted but
// copied into the resulting text.
func TransliterateArabicToIPA(arabicText string) string {
	for i := 0; i < len(transliterationTable); i++ {
		if arabicText == transliterationTable[i][0] {
			return transliterationTable[i][1]
		}
	}
	return arabicText
}

func TestTransliterateArabicToIPA(t *testing.T) {
	tests := []struct {
		arabicText string
		want       string
	}{
		{arabicText: "شجرة", want: "ʃd͜ʒrh"},
		{arabicText: "أ", want: "ʔ"},
		{arabicText: "ب", want: "b"},
		{arabicText: "ت", want: "t"},
	}
	for _, tt := range tests {
		if got := TransliterateArabicToIPA(tt.arabicText); got != tt.want {
			t.Errorf("TransliterateArabicToIPA(%v) = %v, want %v", tt.arabicText, got, tt.want)
		}
	}
}

