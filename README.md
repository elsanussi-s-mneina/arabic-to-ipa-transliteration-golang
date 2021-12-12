# Arabic to IPA command line utility
Converts Arabic script to the International Phonetic Alphabet.

## Supported so far
- consonants only

## Requirements for developers
- Go 1.17.3 (probably works with earlier also)

## Requirements for users:
- none

## How to use:
From the command line enter:

go run main.go < input.txt > output.txt

Where input.txt is an input text file containing Arabic script text.
Output.txt will be created and will contain the transliteration of the Arabic script text.


## To run tests:
Enter the following command:
go test ./...

## To create executable
go build