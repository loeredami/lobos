package main

import "fmt"

func main() {
	var text string = read_file("./examples/print.lobos")
	var tokens []Token = []Token{}
	var lexer Lexer = make_lexer(text)
	tokens = lexer.lex()
	var parser Parser = make_parser(tokens)
	var res ParseResult = parser.parse()

	if res.err == "" {
		fmt.Println("No errors :3")
	} else {
		fmt.Printf("Error: %s\n", res.err)
	}
}
