package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

/*
 * The following are the token types that will be supported in the language.
 */

const (
	ILLEGAL   = "ILLEGAL" // used for marking unknown characters
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMA      = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)
