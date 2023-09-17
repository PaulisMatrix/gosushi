package tinycompiler

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// go cant understand end of char string \0. Represent it in hex form
const EOC string = "\x00"

func StartHere() {
	fmt.Println("Teeny Tiny Compiler")

	//for l.Peek() != EOC {
	//	fmt.Println(l.curChar)
	//	l.NextChar()
	//}
	//token := l.GetToken()
	//for token.tokenKind != EOF {
	//	fmt.Println(string(token.tokenKind))
	//	token = l.GetToken()
	//}

	filePath := flag.String("filePath", "", "File from which to read instructions. Leave empty to read from stdin")
	flag.Parse()

	path, err := filepath.Abs(*filePath)
	if err != nil {
		panic(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	//init lexer
	l := &Lexer{
		source:  string(content),
		curChar: "",
		curPos:  -1,
	}
	//init first character
	l.NextChar()

	//init emitter
	e := &Emitter{
		fullPath: "./output.c",
		code:     "",
		header:   "",
	}

	//init parser
	p := &Parser{
		lexer:     l,
		emitter:   e,
		curToken:  Token{},
		peekToken: Token{},
	}
	//call twice to init current token and peek token
	p.NextToken()
	p.NextToken()

	//start the parser
	p.Program()

	//write to equivalent C code
	e.WriteToFile()

	fmt.Println("Compilation complete!!!")

}
