package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Interface interface {
	TransformRune(pos int, char rune)
	GetValueAsRuneSlice() []rune
}

type Input struct {
	pos  int
	text string
}

var counter string
var b bytes.Buffer

func (i *Input) TransformRune(index int, chars rune) {
	if (index+1)%i.pos == 0 {
		counter = strings.ToUpper(string(chars))
		b.WriteString(counter)
	} else {
		counter = strings.ToLower(string(chars))
		b.WriteString(counter)
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

func main() {
	// s := NewSkipString(3, "Aspiration.com")
	s := NewSkipString(2, "LavanyaBotcha")
	MapString(&s)
	fmt.Println(s)
}

func (i Input) String() string {
	return fmt.Sprintf("Assesement: Capitalize nth element of a given string\nInput Text : %s\nPosition : %d\nOutput Text: %s", i.text, i.pos, b.String())
}
