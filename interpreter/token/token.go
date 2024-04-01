package token

type TokenType string

type Token struct {
	Type		TokenType
	Literal	string
}

var keywords = map[string]TokenType{
	"fn": 		FUNCTION,
	"let": 		LET,
	"const":	CONST,
	"return": RETURN,
	"int":		INT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}


const (
	ILLEGAL = "ILLEGAL"
	EOF			= "EOF"

	// identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y, etc
	INTEGER		= "INTEGER"		// 8675309

	// types
	INT		= "INT"

	// operators
	PLUS = "+"

	// delimiters
	SEMICOLON = ";"
	COMMA			= ","
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"
	COLON			= ":"

	// keywords
	FUNCTION	= "FUNCTION"
	LET				= "LET"
	CONST			= "CONST"
	RETURN		= "RETURN"
)
