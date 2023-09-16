package tinycompiler

import "fmt"

// go cant understand end of char string \0. Represent it in hex form
const EOC string = "\x00"

func StartHere() {
	fmt.Println("lexing first")
	l := &Lexer{
		source:  "IF+-123 f12*THEN/",
		curChar: "",
		curPos:  -1,
	}
	//init first character
	l.NextChar()

	//for l.Peek() != EOC {
	//	fmt.Println(l.curChar)
	//	l.NextChar()
	//}
	token := l.GetToken()
	for token.tokenKind != EOF {
		fmt.Println(string(token.tokenKind))
		token = l.GetToken()
	}
}
