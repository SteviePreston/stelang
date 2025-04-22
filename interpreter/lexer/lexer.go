package lexer

import "github/steviepreston/stelang/interpeter/token"

type Lexer struct {
	input        string
	postiton     int
	readPostiton int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.postiton >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPostiton]
	}
	l.postiton = l.readPostiton
	l.readPostiton += 1
}
