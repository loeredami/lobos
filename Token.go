package main

import "fmt"

const (
	ModeSelect = iota
	Identifier = iota
	Colon      = iota
	Number     = iota
	String     = iota
	SemiColon  = iota
	EOF        = iota
)

type Token struct {
	Type  int
	value string
}

func (token Token) print() {
	fmt.Printf("<Token | Type: %v | Value: %v>\n", token.Type, token.value)
}
