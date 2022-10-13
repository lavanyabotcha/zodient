package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Interface interface {
	TransformRune(pos int, char rune)
	GetValueAsRuneSlice() []rune
}

type Input struct {
	pos  int
	text string
}

var chr string
var b bytes.Buffer

var cntr = 1

func (i *Input) TransformRune(index int, chars rune) {
	if cntr%i.pos == 0 {
		chr = strings.ToUpper(string(chars))
		b.WriteString(chr)
	} else {
		chr = strings.ToLower(string(chars))
		b.WriteString(chr)
	}
	if unicode.IsLetter(chars) || unicode.IsNumber(chars) {
		cntr++
	}
}

func (i *Input) GetValueAsRuneSlice() []rune {
	return []rune(i.text)
}

func MapString(i Interface) {
	for index, value := range i.GetValueAsRuneSlice() {
		i.TransformRune(index, value)
	}
}

func NewSkipString(pos int, text string) Input {

	input := Input{pos: pos, text: text}
	return input
}

func (i Input) String() string {
	return fmt.Sprintf("Assesement: Capitalize nth element of a given string\nInput Text : %s\nPosition : %d\nOutput Text: %s", i.text, i.pos, b.String())
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}
