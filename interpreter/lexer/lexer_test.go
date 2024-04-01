/*
 * TODO: we'll need to also attach filenames and line numbers
 * to tokens to better track down lexing and parsing errors.
 * To do that we might need to initialize the lexer with an io.Reader
 * and the filename. For now, we can skip this, but it's definitely
 * going to need to be done eventually.
 *
*/

package lexer

import (
	"testing"
	"interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `+;,(){}:`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COLON, ":"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, tt.expectedLiteral, tok.Literal)
		}
	}
}
