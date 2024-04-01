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
	"alpha/token"
)

func TestNextToken(t *testing.T) {
	input := `let five: int { 5 };
	let ten: int { 10 };
	
	fn add(x, y): int {
		return x + y;
	};
	
	let result: int { add(five, ten); };
	!-/*5;
	5 < 10 > 5;
	if else;
	const true false;
	== !=;
	<= >=;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.COLON, ":"},
		{token.INT, "int"},
		{token.LBRACE, "{"},
		{token.INTEGER, "5"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.COLON, ":"},
		{token.INT, "int"},
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
		{token.INT, "int"},
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
		{token.INT, "int"},
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
		{token.CONST, "const"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.IEQ, "=="},
		{token.NEQ, "!="},
		{token.SEMICOLON, ";"},
		{token.LEQ, "<="},
		{token.GEQ, ">="},
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
