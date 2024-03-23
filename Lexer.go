package main

import (
	"math/big"
	"strings"
)

const letters string = "abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits string = "0123456789"

var letters_and_digits string = string(append([]rune(letters), []rune(digits)...))

type Lexer struct {
	text     string
	tokens   []Token
	done     bool
	idx      uint64
	cur_char rune
}

func (lexer *Lexer) collect_mode() {
	lexer.advance()
	var mode []rune = []rune{}
	for {
		if !strings.ContainsRune(letters, lexer.cur_char) {
			break
		}
		mode = append(mode, lexer.cur_char)
		lexer.advance()
	}
	lexer.tokens = append(lexer.tokens, Token{ModeSelect, string(mode)})
	lexer.advance()
}

func (lexer *Lexer) collect_identifier() {
	var identifier []rune = []rune{}
	for {
		if !strings.ContainsRune(letters_and_digits, lexer.cur_char) {
			break
		}
		identifier = append(identifier, lexer.cur_char)
		lexer.advance()
	}
	lexer.tokens = append(lexer.tokens, Token{Identifier, string(identifier)})
}

func (lexer *Lexer) add_colon() {
	lexer.advance()
	lexer.tokens = append(lexer.tokens, Token{Colon, ""})
}

func (lexer *Lexer) add_semi_colon() {
	lexer.advance()
	lexer.tokens = append(lexer.tokens, Token{SemiColon, ""})
}

func (lexer *Lexer) collect_number() {
	var num []rune = []rune{}
	var is_hex bool = false
	var at_start bool = true
	for {
		if !(strings.ContainsRune(digits, lexer.cur_char) || (is_hex && strings.ContainsRune(digits+"ABCDEFabcdef", lexer.cur_char))) {
			break
		}
		if at_start {
			if lexer.cur_char == '0' {
				lexer.advance()
				if lexer.cur_char == 'x' {
					is_hex = true
					lexer.advance()
				}
			}
			at_start = false
		}
		num = append(num, lexer.cur_char)
		lexer.advance()
	}
	if is_hex {
		n := new(big.Int)
		n.SetString(string(num), 16)
		var val int64 = n.Int64()
		lexer.tokens = append(lexer.tokens, Token{Number, string(rune(val))})
		return
	}
	lexer.tokens = append(lexer.tokens, Token{Number, string(num)})
}

func (lexer *Lexer) collect_string() {
	var str []rune = []rune{}

	lexer.advance()

	for {
		if lexer.cur_char == '"' {
			break
		}
		str = append(str, lexer.cur_char)
		lexer.advance()
	}

	lexer.advance()

	lexer.tokens = append(lexer.tokens, Token{String, string(str)})
}

func (lexer *Lexer) update_char() {
	if lexer.idx >= uint64(len(lexer.text)) {
		lexer.done = true
		return
	}
	lexer.cur_char = []rune(lexer.text)[lexer.idx]
}

func (lexer *Lexer) advance() {
	lexer.idx++
	lexer.update_char()
}

func (lexer *Lexer) skip_line() {
	for {
		if lexer.cur_char == '\n' {
			break
		}
		lexer.advance()
	}
	lexer.advance()
}

func (lexer *Lexer) lex() []Token {
	lexer.update_char()
	for {
		if lexer.done {
			break
		}

		switch {
		case lexer.cur_char == '[':
			lexer.collect_mode()
		case strings.ContainsRune(letters, lexer.cur_char):
			lexer.collect_identifier()
		case strings.ContainsRune(digits, lexer.cur_char):
			lexer.collect_number()
		case lexer.cur_char == ':':
			lexer.add_colon()
		case lexer.cur_char == ';':
			lexer.add_semi_colon()
		case lexer.cur_char == '"':
			lexer.collect_string()
		case lexer.cur_char == '#':
			lexer.skip_line()
		default:
			lexer.advance()
		}
	}

	lexer.tokens = append(lexer.tokens, Token{EOF, ""})

	return lexer.tokens
}
