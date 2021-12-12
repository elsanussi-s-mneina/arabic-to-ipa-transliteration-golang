package transliteration_test

import (
	"testing"

	"youtube.com.fake/GrassDewMorning/arabicToIPA/transliteration"
)

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
		{arabicText: "مرحبا", want: "mrħbaː"},
		
	}
	for _, tt := range tests {
		if got := transliteration.TransliterateArabicToIPA(tt.arabicText); got != tt.want {
			t.Errorf("TransliterateArabicToIPA(%v) = %v, want %v", tt.arabicText, got, tt.want)
		}
	}
}
