package tinycompiler

import (
	"fmt"
	"os"
	"unicode"
)

type TokenKind string

const (
	EOF     TokenKind = "EOF"
	NEWLINE TokenKind = "NEWLINE"
	NUMBER  TokenKind = "NUMBER"
	IDENT   TokenKind = "IDENT"
	STRING  TokenKind = "STRING"
	//keywords
	LABLE    TokenKind = "LABEL"
	GOTO     TokenKind = "GOTO"
	PRINT    TokenKind = "PRINT"
	INPUT    TokenKind = "INPUT"
	LET      TokenKind = "LET"
	IF       TokenKind = "IF"
	THEN     TokenKind = "THEN"
	ENDIF    TokenKind = "ENDIF"
	WHILE    TokenKind = "WHILE"
	REPEAT   TokenKind = "REPEAT"
	ENDWHILE TokenKind = "ENDWHILE"
	//operators
	EQ       TokenKind = "EQUALTO"
	PLUS     TokenKind = "PLUS"
	MINUS    TokenKind = "MINUS"
	ASTERISK TokenKind = "ASTERISK"
	SLASH    TokenKind = "SLASH"
	EQEQ     TokenKind = "EQUALTOEQUALTO"
	NOTEQ    TokenKind = "NOTEQUALTO"
	LT       TokenKind = "LESSTHAN"
	LTEQ     TokenKind = "LESSTHANEQUALTO"
	GT       TokenKind = "GREATERTHAN"
	GTEQ     TokenKind = "GREATERTHANEQUALTO"
)

var KeyWordList = []TokenKind{
	LABLE, GOTO, PRINT, INPUT, LET, IF, THEN, ENDIF, WHILE, REPEAT, ENDWHILE,
}

type LexerIface interface {
	//process the next character
	NextChar()
	//return the lookahead character
	Peek() string
	//invalid token found, print error message and exit
	Abort(string)
	//skip whitespace except newlines, which we will use to indicate the end of a statement
	SkipWhiteSpace()
	//skip comments in the code
	SkipComment()
	//return the next token
	GetToken() Token
	//check for special characters
	IsSpecialCharacter(char string) bool
}

var _ LexerIface = (*Lexer)(nil)

type Lexer struct {
	source  string //source or the current string line to the lexer
	curChar string //current character in the string
	curPos  int    //current position in the string
}

type TokenIface interface {
	IsKeyWord(keyword TokenKind) bool
}

var _ TokenIface = (*Token)(nil)

type Token struct {
	tokenText string
	tokenKind TokenKind
}

func (t *Token) IsKeyWord(keyword TokenKind) bool {
	for _, kind := range KeyWordList {
		if keyword == kind {
			return true
		}
	}
	return false
}

func (l *Lexer) NextChar() {
	l.curPos++
	if l.curPos >= len(l.source) {
		l.curChar = EOC
	} else {
		l.curChar = string(l.source[l.curPos])
	}

}
func (l *Lexer) Peek() string {
	if (l.curPos + 1) >= len(l.source) {
		return EOC
	}
	return string(l.source[l.curPos+1])
}

func (l *Lexer) Abort(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func (l *Lexer) SkipWhiteSpace() {
	for l.curChar == " " || l.curChar == "\t" || l.curChar == "\r" {
		l.NextChar()
	}
}

func (l *Lexer) SkipComment() {
	// all text after "#" are skipped except the newline

	if l.curChar == "#" {
		for l.curChar != "\n" {
			l.NextChar()
		}
	}
}

func (l *Lexer) IsSpecialCharacter(char string) bool {
	if char == "\r" || char == "\n" || char == "\t" || char == "\\" || l.curChar == "%" {
		return true
	}
	return false
}
func (l *Lexer) GetToken() Token {
	// check for characters except newline
	l.SkipWhiteSpace()
	//check for comments
	l.SkipComment()
	token := Token{}
	switch {
	case l.curChar == "+":
		token = Token{
			tokenText: l.curChar,
			tokenKind: PLUS,
		}
	case l.curChar == "-":
		token = Token{
			tokenText: "-",
			tokenKind: MINUS,
		}
	case l.curChar == "*":
		token = Token{
			tokenText: "*",
			tokenKind: ASTERISK,
		}
	case l.curChar == "/":
		token = Token{
			tokenText: "/",
			tokenKind: SLASH,
		}
	case l.curChar == "\n":
		token = Token{
			tokenText: "\n",
			tokenKind: NEWLINE,
		}
	case l.curChar == "=":
		// = or  ==
		if l.Peek() == "=" {
			lastChar := l.curChar
			l.NextChar()
			token = Token{
				tokenText: lastChar + l.curChar,
				tokenKind: EQEQ,
			}
		} else {
			token = Token{
				tokenText: l.curChar,
				tokenKind: EQ,
			}
		}
	case l.curChar == ">":
		//> or >=
		if l.Peek() == "=" {
			lastChar := l.curChar
			l.NextChar()
			token = Token{
				tokenText: lastChar + l.curChar,
				tokenKind: GTEQ,
			}
		} else {
			token = Token{
				tokenText: l.curChar,
				tokenKind: GT,
			}
		}
	case l.curChar == "<":
		//< or <=
		if l.Peek() == "=" {
			lastChar := l.curChar
			l.NextChar()
			token = Token{
				tokenText: lastChar + l.curChar,
				tokenKind: LTEQ,
			}
		} else {
			token = Token{
				tokenText: l.curChar,
				tokenKind: LT,
			}
		}
	case l.curChar == "!":
		if l.Peek() == "=" {
			lastChar := l.curChar
			l.NextChar()
			token = Token{
				tokenText: lastChar + l.curChar,
				tokenKind: NOTEQ,
			}
		} else {
			msg := fmt.Sprintf("Expected !=, got ! and %s", l.Peek())
			l.Abort(msg)
		}
	case l.curChar == "\"":
		l.NextChar()
		startPos := l.curPos

		// " \"This is a String\" "
		for l.curChar != "\"" {
			// dont allow special characters in the string
			if l.IsSpecialCharacter(l.curChar) {
				msg := fmt.Sprintf("illegal character in the string: %s", l.curChar)
				l.Abort(msg)
			}
			l.NextChar()
		}
		text := l.source[startPos:l.curPos] //get the string content
		token = Token{
			tokenText: text,
			tokenKind: STRING,
		}
	case unicode.IsDigit(rune(l.curChar[0])) == true:
		//get all consecutive digits and decimal too if any
		startPos := l.curPos
		for unicode.IsDigit(rune(l.Peek()[0])) {
			l.NextChar()
		}

		//decimal
		if l.Peek() == "." {
			l.NextChar()

			//must have at least one digit after decimal point
			if !unicode.IsDigit(rune(l.Peek()[0])) {
				msg := fmt.Sprintf("illegal character in number: %s", l.Peek())
				l.Abort(msg)
			}
			for unicode.IsDigit(rune(l.Peek()[0])) {
				l.NextChar()
			}
		}
		text := l.source[startPos:l.curPos]
		token = Token{
			tokenText: text,
			tokenKind: NUMBER,
		}
	case unicode.IsLetter(rune(l.curChar[0])) == true:
		//check if it is an indentifier or a keyword
		startPos := l.curPos

		for unicode.IsLetter(rune(l.Peek()[0])) || unicode.IsDigit(rune(l.Peek()[0])) {
			l.NextChar()
		}
		text := l.source[startPos : l.curPos+1]
		token = Token{
			tokenText: text,
		}
		// its a keyword
		if token.IsKeyWord(TokenKind(token.tokenText)) {
			token.tokenKind = TokenKind(token.tokenText)
		} else {
			token.tokenKind = IDENT
		}
	case l.curChar == EOC:
		token = Token{
			tokenText: EOC,
			tokenKind: EOF,
		}
	default:
		// invalid token
		msg := fmt.Sprintf("Lexing error. Unkown token: %s", l.curChar)
		l.Abort(msg)

	}
	l.NextChar()
	return token
}
