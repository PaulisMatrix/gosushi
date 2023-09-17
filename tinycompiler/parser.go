package tinycompiler

import (
	"fmt"
	"os"
)

type ParserIface interface {
	// return true if the current token matches
	CheckToken(TokenKind) bool
	//return true if the next token matches
	CheckPeek(TokenKind) bool
	//try to match current token. If not, error. Advances the current token
	Match(TokenKind)
	//advances to the next token
	NextToken()
	//abort the execution if on an error
	Abort(msg string)
	//parser starting point
	Program()
	//check for comparison operator
	IsComparisonOperator() bool

	//parse all the intructions. check readme for grammer.
	//parse statement
	Statement()
	//parse nl
	Nl()
	//parse expression
	Expression()
	//parse term
	Term()
	//parse unary
	Unary()
	//parse primary
	Primary()
	//parse Comparison
	Comparison()
}

var _ ParserIface = (*Parser)(nil)

// parser is used to keep track of the current token and checks if the code matches the grammer.
type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token

	variables      []string //list of all variables declared so far.
	lablesDeclared []string //list of all lables declared so far.
	gotoDeclared   []string //list of all gotos declared so far.

}

func (p *Parser) CheckToken(tokenKind TokenKind) bool {
	return p.curToken.tokenKind == tokenKind
}

func (p *Parser) CheckPeek(tokenKind TokenKind) bool {
	return p.peekToken.tokenKind == tokenKind
}

func (p *Parser) Match(tokenKind TokenKind) {
	// try to match current token. If matches proceed to next one, if not error out.
	if !p.CheckToken(tokenKind) {
		msg := fmt.Sprintf("Expected %s but got %s", string(tokenKind), string(p.curToken.tokenKind))
		p.Abort(msg)
	}
	p.NextToken()
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.GetToken()
}

func (p *Parser) Abort(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func (p *Parser) Program() {
	fmt.Println("PROGRAM")

	//skip newlines at the start of the program
	for p.CheckToken(NEWLINE) {
		p.NextToken()
	}

	//parse all the statements in the program till EOF
	for !p.CheckToken(EOF) {
		p.Statement()
	}

	//check that each label gotoed has been declared or not
	for _, gotos := range p.gotoDeclared {
		// if labelsDeclared is empty then no labels are present so far
		if len(p.lablesDeclared) == 0 {
			msg := fmt.Sprintf("attempting to GOTO to an undeclared label: %s", gotos)
			p.Abort(msg)
		}
		for _, labels := range p.lablesDeclared {
			if gotos != labels {
				msg := fmt.Sprintf("attempting to GOTO to an undeclared label: %s", gotos)
				p.Abort(msg)
			}
		}
	}
}

func (p *Parser) Statement() {
	//check the first token to see what kind of statement it is

	if p.CheckToken(PRINT) { //"PRINT" (expression | string)
		fmt.Println("STATEMENT-PRINT")
		p.NextToken()

		if p.CheckToken(STRING) {
			//string
			p.NextToken()
		} else {
			//expression
			p.Expression()
		}
	} else if p.CheckToken(IF) { //"IF" comparison "THEN" nl {statement} "ENDIF"
		fmt.Println("STATEMENT-IF")
		p.NextToken()
		p.Comparison()

		p.Match(THEN)
		p.Nl()

		for !p.CheckToken(ENDIF) {
			//zero or more statements
			p.Statement()
		}
		p.Match(ENDIF)
	} else if p.CheckToken(WHILE) { //"WHILE" comparison "REPEAT" nl {statement} "ENDWHILE"
		fmt.Println("STATEMENT-WHILE")
		p.NextToken()
		p.Comparison()

		p.Match(REPEAT)
		p.Nl()

		for !p.CheckToken(ENDWHILE) {
			//zero or more statements
			p.Statement()
		}
		p.Match(ENDWHILE)
	} else if p.CheckToken(LABEL) { //"LABEL" ident
		fmt.Println("STATEMENT-LABEL")
		p.NextToken()
		if p.CheckToken(IDENT) {
			//if label found, record it
			found := false
			for _, vars := range p.lablesDeclared {
				if p.curToken.tokenText == vars {
					//already present, dont add again
					found = true
					msg := fmt.Sprintf("label %s exists already", p.curToken.tokenText)
					p.Abort(msg)
					break
				}
			}
			if !found {
				p.lablesDeclared = append(p.lablesDeclared, p.curToken.tokenText)
			}
		} else {
			msg := fmt.Sprintf("expected identifier, got %s", p.curToken.tokenText)
			p.Abort(msg)
		}
		p.Match(IDENT)
	} else if p.CheckToken(GOTO) { //"GOTO" ident
		fmt.Println("STATEMENT-GOTO")
		p.NextToken()
		if p.CheckToken(IDENT) {
			//record all gotos
			p.gotoDeclared = append(p.gotoDeclared, p.curToken.tokenText)
		} else {
			msg := fmt.Sprintf("expected identifier, got %s", p.curToken.tokenText)
			p.Abort(msg)
		}
		p.Match(IDENT)
	} else if p.CheckToken(LET) { //LET" ident "=" expression
		fmt.Println("STATEMENT-LET")
		p.NextToken()

		if p.CheckToken(IDENT) {
			//if identifier(var) found, record it
			found := false
			for _, vars := range p.variables {
				if p.curToken.tokenText == vars {
					//if already exists, skip
					continue
				}
			}
			if !found {
				p.variables = append(p.variables, p.curToken.tokenText)
			}
		} else {
			msg := fmt.Sprintf("expected identifier, got %s", p.curToken.tokenText)
			p.Abort(msg)
		}
		p.Match(IDENT)
		p.Match(EQ)
		p.Expression()
	} else if p.CheckToken(INPUT) { //"INPUT" ident
		fmt.Println("STATEMENT-INPUT")
		p.NextToken()
		if p.CheckToken(IDENT) {
			//if identifier(var) found, record it
			found := false
			for _, vars := range p.variables {
				if p.curToken.tokenText == vars {
					//if already exists, skip
					continue
				}
			}
			if !found {
				p.variables = append(p.variables, p.curToken.tokenText)
			}
		} else {
			msg := fmt.Sprintf("expected identifier, got %s", p.curToken.tokenText)
			p.Abort(msg)
		}
		p.Match(IDENT)
	} else {
		msg := fmt.Sprintf("invalid statement %s", p.curToken.tokenText)
		p.Abort(msg)
	}
	p.Nl()
}

// n1 :: "\n"+
func (p *Parser) Nl() {
	fmt.Println("NEWLINE")

	//ideally just one newline
	p.Match(NEWLINE)

	//but allow for extra new lines too
	for p.CheckToken(NEWLINE) {
		p.NextToken()
	}
}

// expression :: term {( "-" | "+" ) term}
func (p *Parser) Expression() {
	fmt.Println("EXPRESSION")

	p.Term()

	for p.CheckToken(MINUS) || p.CheckToken(PLUS) {
		p.NextToken()
		p.Term()
	}
}

// term ::= unary {( "/" | "*" ) unary}
func (p *Parser) Term() {
	fmt.Println("TERM")
	p.Unary()

	for p.CheckToken(SLASH) || p.CheckToken(ASTERISK) {
		p.NextToken()
		p.Unary()
	}
}

// unary ::= {"+" | "-"} primary
func (p *Parser) Unary() {
	fmt.Println("UNARY")

	for p.CheckToken(PLUS) || p.CheckToken(MINUS) {
		p.NextToken()
	}
	p.Primary()
}

// primary ::= number | ident
func (p *Parser) Primary() {
	fmt.Println("PRIMARY (" + p.curToken.tokenText + ")")

	if p.CheckToken(NUMBER) {
		p.NextToken()
	} else if p.CheckToken(IDENT) {
		foundVar := false
		//ensure variable you accessing, exists first
		for _, vars := range p.variables {
			if p.curToken.tokenText == vars {
				p.NextToken()
				foundVar = true
				break
			}
		}
		// didn't find
		if !foundVar {
			msg := fmt.Sprintf("referencing variable before assignment: %s", p.curToken.tokenText)
			p.Abort(msg)
		}
	} else {
		msg := fmt.Sprintf("Unexpected token: %s", p.curToken.tokenText)
		p.Abort(msg)
	}

}

func (p *Parser) IsComparisonOperator() bool {
	return p.CheckToken(EQEQ) || p.CheckToken(NOTEQ) || p.CheckToken(GT) || p.CheckToken(GTEQ) || p.CheckToken(LT) || p.CheckToken(LTEQ)
}

// comparison ::= expression (("==" | "!=" | ">" | ">=" | "<" | "<=") expression)+
func (p *Parser) Comparison() {
	p.Expression()

	// there should be at least one operator
	if p.IsComparisonOperator() {
		p.NextToken()
		p.Expression()

	} else {
		msg := fmt.Sprintf("invalid token %s , expected an operator", p.curToken.tokenText)
		p.Abort(msg)
	}

	//could be more than one operator
	for p.IsComparisonOperator() {
		p.NextToken()
		p.Expression()
	}

}
