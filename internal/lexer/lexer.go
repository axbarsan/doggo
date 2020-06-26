package lexer

import (
	"github.com/axbarsan/doggo/internal/token"
)

type Lexer struct {
	input string
	// position is the last read position.
	position int
	// readPosition is the position that we're gonna read from next.
	readPosition int
	// ch is the current char under examination.
	ch byte
}

// The lexer will parse the source code and extract known tokens, which will be later turned into the AST of the program.
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token

	// TODO: After converting types to integers, use 'iota' to categorize token types and parse this easier.
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch))
	case '/':
		tok = newToken(token.SLASH, string(l.ch))
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readWithValidator(isLetter)
			tok.Type = token.LookupIdentifier(tok.Literal)

			// Return early because we already read the next char in the 'readIdentifier' method.
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readWithValidator(isDigit)
			tok.Type = token.INT

			return tok
		}

		tok = newToken(token.ILLEGAL, string(l.ch))
	}

	l.readChar()

	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition < len(l.input) {
		return l.input[l.readPosition] // Current character.
	}

	return 0 // ASCII code for the 'NUL' character, meaning: Nothing read yet, or in this case: EOF
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readWithValidator(v func(c byte) bool) string {
	pos := l.position

	for v(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) makeTwoCharToken(t token.Type) token.Token {
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)

	tok := newToken(t, literal)

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.Type, ch string) token.Token {
	t := token.Token{
		Type:    tokenType,
		Literal: ch,
	}

	return t
}

func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
