package lexer

import (
	"testing"

	"github.com/patricklizon/interpreter-in-go/token"
)

func TestNextToken(t *testing.T) {
	input := `=:;(){},+`

	tests := []struct {
		expectedTokenType token.TokenType
		expectedLiteral   string
	}{
		{token.ASSIGN, "="},
		{token.COLON, ":"},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMA, ","},
		{token.PLUS, "+"},
		{token.EOF, ""},
	}

	l := New(input)

	for idx, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", idx, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tt.expectedLiteral, tok.Literal)
		}
	}
}
