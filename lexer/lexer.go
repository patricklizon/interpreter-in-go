package lexer

import "github.com/patricklizon/interpreter-in-go/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input,
		// it's not necessary to initialize position and readPosition to 0,
		// as they are zero-valued by default in Go
		// but as I am learning Go, I will initialize them anyway :)
		readPosition: 0,
	}

	// sets the current character to the first character in the input string
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	switch l.ch {
	case '=':
		tok = makeToken(token.ASSIGN, l.ch)
	case ':':
		tok = makeToken(token.COLON, l.ch)
	case ';':
		tok = makeToken(token.SEMICOLON, l.ch)
	case '(':
		tok = makeToken(token.LPAREN, l.ch)
	case ')':
		tok = makeToken(token.RPAREN, l.ch)
	case ',':
		tok = makeToken(token.COMA, l.ch)
	case '+':
		tok = makeToken(token.PLUS, l.ch)
	case '{':
		tok = makeToken(token.LBRACE, l.ch)
	case '}':
		tok = makeToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.readChar()

	return tok
}

// sets the next character and advances our position in the input string
func (l *Lexer) readChar() {
	l.ch = func() byte {
		if l.readPosition >= len((l.input)) {
			// ASCII code for "NUL", which means "end of input"
			return 0
		} else {
			return l.input[l.readPosition]
		}
	}()
	l.position = l.readPosition
	l.readPosition += 1
}

// makes token based on the current character
func makeToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
