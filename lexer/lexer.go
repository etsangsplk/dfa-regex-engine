// Package lexer implements lexer for a simple regular expression.
package lexer

import (
	"github.com/8ayac/dfa-regex-engine/token"
)

// Lexer has a slice of symbols to analyze.
type Lexer struct {
	s []rune // string to be analyzed
}

// NewLexer returns a new Lexer.
// This constructor create a sequence of symbols from
// the string given in the argument and hold it.
func NewLexer(s string) *Lexer {
	return &Lexer{
		s: []rune(s),
	}
}

// getToken return the appropriate token for the given symbol.
func (l *Lexer) getToken(r rune) *token.Token {
	switch r {
	case '\x00':
		return token.NewToken(r, token.EOF)
	case '\\':
		return token.NewToken(r, token.CHARACTER)
	case '|':
		return token.NewToken(r, token.OpeUnion)
	case '(':
		return token.NewToken(r, token.LPAREN)
	case ')':
		return token.NewToken(r, token.RPAREN)
	case '*':
		return token.NewToken(r, token.OpeStar)
	default:
		return token.NewToken(r, token.CHARACTER)
	}
}

// Scan returns the all token to which converted from
// the symbol slice held in Lexer struct.
func (l *Lexer) Scan() (tokenList []*token.Token) {
	for _, r := range l.s {
		if len(l.s) == 0 {
			tokenList = append(tokenList, token.NewToken('\x00', token.EOF))
		}
		tokenList = append(tokenList, l.getToken(r))
	}
	return
}
