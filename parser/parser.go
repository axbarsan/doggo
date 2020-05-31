package parser

import (
	"fmt"

	"github.com/axbarsan/doggo/ast"
	"github.com/axbarsan/doggo/lexer"
	"github.com/axbarsan/doggo/token"
)

// The parser takes the input text and builds a hierarchical data structure. (AST)
// It gives a structural representation of the input, checking the correct syntax
// in the process. Parsing is also called 'Syntactic analysis'.
//
// 'Abstract' in AST comes from the fact that certain details from the source code
// are omitted in the AST

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read 2 tokens, so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()

		return true
	}

	p.peekError(t)

	return false
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.CONST:
		return p.parseConstStatement()
	default:
		return nil
	}
}

func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(tok token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tok, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
