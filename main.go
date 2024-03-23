package main

func main() {
	var text string = read_file("./examples/print.lobos")
	var tokens []Token = []Token{}
	var lexer Lexer = Lexer{text, tokens, false, 0, ' '}

	tokens = lexer.lex()

	for i := 0; i < len(tokens); i++ {
		tokens[i].print()
	}
}
