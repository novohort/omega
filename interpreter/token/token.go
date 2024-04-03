package token

type TokenType string

type Token struct {
	Type		TokenType
	Literal	string
}

var keywords = map[string]TokenType{
	"fn": 		FUNCTION,
	"let": 		LET,
	"return": RETURN,
	"true":		TRUE,
	"false":	FALSE,
	"if":			IF,
	"else":		ELSE,
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
	IDENT = "IDENT"

	// data types
	INTEGER		= "INTEGER"		// 8675309
	STRING		= "STRING"		// "zaphod beeblebrox"

	// comparisons
	IEQ	= "=="
	NEQ = "!="
	LEQ = "<="
	GEQ = ">="

	// operators
	PLUS 			= "+"
	ASSIGN 		= "="
	MINUS 		= "-"
	BANG			= "!"
	ASTERISK	= "*"
	FSLASH		= "/"
	LT				= "<"
	GT				= ">"

	// delimiters
	SEMICOLON = ";"
	COMMA			= ","
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"
	COLON			= ":"
	LBRACKET	= "["
	RBRACKET	=	"]"

	// keywords
	FUNCTION	= "FUNCTION"
	LET				= "LET"
	RETURN		= "RETURN"
	TRUE			= "TRUE"
	FALSE			= "FALSE"
	IF				= "IF"
	ELSE			= "ELSE"
)
