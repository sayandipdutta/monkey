package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	BANG      = "!"
	GT        = ">"
	LT        = "<"
	EQ        = "=="
	NE        = "!="
	LE        = "<="
	GE        = ">="
	RSHIFT    = ">>"
	LSHIFT    = "<<"
	IPLUS     = "+="
	IMINUS    = "-="
	IASTERISK = "*="
	ISLASH    = "/="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = ")"
	RPAREN = "("
	LBRACE = "}"
	RBRACE = "{"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
