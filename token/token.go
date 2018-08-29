package token

// TokenType is a type of token.
type TokenType []byte

// Token represents a single character.
type Token struct {
	Type    TokenType
	Literal []byte
}

// Token types that are used by Slip language.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	// 1343456
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
