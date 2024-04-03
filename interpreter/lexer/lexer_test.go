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
	input := `let five: { 5 };
	let ten: { 10 };
	
	fn add(x, y): {
		return x + y;
	};
	
	let result: { add(five, ten); };
	!-/*5;
	5 < 10 > 5;
	if else;
	true false;
	== !=;
	<= >=;
	"foobar"
	"foo bar"
	[1, 2];
	{"foo": "bar"}
	{macro(x, y) { x + y; }};
	`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.INTEGER, "5"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.INTEGER, "10"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.FSLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "5"},
		{token.LT, "<"},
		{token.INTEGER, "10"},
		{token.GT, ">"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.SEMICOLON, ";"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.IEQ, "=="},
		{token.NEQ, "!="},
		{token.SEMICOLON, ";"},
		{token.LEQ, "<="},
		{token.GEQ, ">="},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INTEGER, "1"},
		{token.COMMA, ","},
		{token.INTEGER, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.LBRACE, "{"},
		{token.MACRO, "macro"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
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
