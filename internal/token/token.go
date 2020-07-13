package token

// These are the tokens that our lexer will extract from the source code.

type Type string

type Token struct {
	// Type represents the type of the token (e.g. FUNCTION).
	Type Type
	// Literal represents the token, but in a literal fashion (e.g. func).
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // An unknown token.
	EOF     = "EOF"     // End of file.

	// Identifiers and literals.
	IDENT  = "IDENT"  // Function / variable name.
	INT    = "INT"    // Integer.
	STRING = "STRING" // String.

	// Operators.
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords.
	FUNCTION = "FUNCTION"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var (
	keywords = map[string]Type{
		"fn":     FUNCTION,
		"const":  CONST,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
	}
)

func LookupIdentifier(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}
