package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/toke"
)

type Parser struct {
	l *lexer.Token

	curToken  token.Token
	peekToken token.Token
}
