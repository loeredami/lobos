package main

import "math/big"

const (
	DATA_M     = iota
	REGISTER_M = iota
	EXE_M      = iota
	NO_M       = iota
)

const (
	CACHE_R = iota
	MEM_R   = iota
	KEY_R   = iota
	MOUS_R  = iota
	DISK_R  = iota
)

type Statement struct {
	command   string
	arguments []uint64
}

type ParseResult struct {
	statements   []Statement
	register_map map[string]Register
	err          string
}

type Register struct {
	idx   uint64
	reg_t int
}

type Parser struct {
	// Unprocessed
	tokens []Token

	done bool

	// maps
	label_map    map[string]uint64
	var_map      map[string]uint64
	register_map map[string]Register

	mode int

	idx int

	cur_token Token

	parse_result ParseResult
}

func make_parser(tokens []Token) Parser {
	var p Parser = Parser{tokens, false, make(map[string]uint64), make(map[string]uint64), make(map[string]Register), NO_M, 0, Token{NT, ""}, ParseResult{[]Statement{}, make(map[string]Register), ""}}
	p.update_token()
	return p
}

func (parser *Parser) update_token() {
	if parser.idx > len(parser.tokens) {
		parser.done = true
		return
	}
	parser.cur_token = parser.tokens[parser.idx]
}

func (parser *Parser) advance() {
	parser.idx++
	parser.update_token()
}

func (parser *Parser) no_mode() {
	for {
		if parser.cur_token.Type == ModeSelect {
			if parser.cur_token.value == "DATA"{
				parser.mode = DATA_M
				break
			} else if parser.cur_token.value == "REGISTERS" {
				parser.mode = REGISTER_M
				break
			} else if parser.cur_token.value == "EXECUTABLE" {
				parser.mode = EXE_M
				break
			} else {
				parser.mode = NO_M
				break
			}
		} else {
			parser.advance()
		}
	}
}

func (parser *Parser) make_result() ParseResult {
	parser.parse_result.err = ""
	parser.parse_result.register_map = parser.register_map
	return parser.parse_result
}

func (parser *Parser) arr_read() []uint64 {
	var result []uint64 = []uint64{}
	return result
}

func (parser *Parser) add_statement(command string, args []uint64) {
	parser.parse_result.statements = append(parser.parse_result.statements, Statement{command, args})
}

func (parser *Parser) add_memory(value uint64) {
	// no implemented
}

func (parser *Parser) add_data_entry() {
	parser.advance()
	for {
		if parser.cur_token.Type == SemiColon {
			break
		}
		if parser.cur_token.Type == ModeSelect {
			parser.mode = NO_M
			return
		}
		if parser.cur_token.Type == String {
			for idx := 0; idx < len(parser.cur_token.value); idx++ {
				parser.add_memory(uint64([]rune(parser.cur_token.value)[idx]))
			} 
			parser.advance()
		} else {
			n := new(big.Int)
			n.SetString(parser.cur_token.value, 10)
			parser.add_memory(n.Uint64())
			parser.advance()
		}
	}
	parser.advance()
}

func (parser *Parser) map_data() {
	if parser.cur_token.Type == Identifier {
		parser.add_data_entry()
	} else {
		parser.advance()
	}
}

func (parser *Parser) map_registers() {

}

func (parser *Parser) make_exec() {

}

func (parser *Parser) parse() ParseResult {
	for {
		if parser.done {
			break
		}
		if parser.mode == NO_M {
			parser.no_mode()
		} else if parser.mode == DATA_M {
			parser.map_data()
		} else if parser.mode == REGISTER_M {
			parser.map_registers()
		} else if parser.mode == EXE_M {
			parser.make_exec()
		} else {
			parser.advance()
		}
	}
	return parser.make_result()
}
