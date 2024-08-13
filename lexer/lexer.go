package lexer

import (
	"github.com/sayandipdutta/monkey/token"
)

type Lexer struct {
	input        string
	currPosition int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (lexer *Lexer) readChar() {
	if lexer.nextPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.nextPosition]
	}
	lexer.currPosition = lexer.nextPosition
	lexer.nextPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhiteSpace()
	currChar := string(lexer.ch)

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.EQ, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.ASSIGN, currChar)
		}
	case '+':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.IPLUS, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.PLUS, currChar)
		}
	case '-':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.IMINUS, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.MINUS, currChar)
		}
	case '*':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.IASTERISK, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.ASTERISK, currChar)
		}
	case '/':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.ISLASH, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.SLASH, currChar)
		}
	case '>':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.GE, currChar+string(lexer.ch))
		} else if lexer.peekChar() == '>' {
			lexer.readChar()
			tok = newToken(token.RSHIFT, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.GT, currChar)
		}
	case '<':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.LE, currChar+string(lexer.ch))
		} else if lexer.peekChar() == '<' {
			lexer.readChar()
			tok = newToken(token.LSHIFT, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.LT, string(lexer.ch))
		}
	case '!':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = newToken(token.NE, currChar+string(lexer.ch))
		} else {
			tok = newToken(token.BANG, string(lexer.ch))
		}
	case '(':
		tok = newToken(token.LPAREN, string(lexer.ch))
	case ')':
		tok = newToken(token.RPAREN, string(lexer.ch))
	case '{':
		tok = newToken(token.LBRACE, string(lexer.ch))
	case '}':
		tok = newToken(token.RBRACE, string(lexer.ch))
	case ',':
		tok = newToken(token.COMMA, string(lexer.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(lexer.ch))
	case 0:
		tok = newToken(token.EOF, "EOF")
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, "ILLEGAL")
		}
	}
	lexer.readChar()
	return tok
}

func newToken(toktype token.TokenType, literal string) token.Token {
	return token.Token{Type: toktype, Literal: literal}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.currPosition
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[startPosition:lexer.currPosition]
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.currPosition
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	got := lexer.input[startPosition:lexer.currPosition]
	return got
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\r' || lexer.ch == '\n' {
		lexer.readChar()
	}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.nextPosition == 0 {
		return 0
	} else {
		return lexer.input[lexer.nextPosition]
	}
}
