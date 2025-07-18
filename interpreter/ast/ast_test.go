package ast

import (
	token "github.com/steviepreston/stelang/interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifer{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifer{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong, Got: %q", program.String())
	}
}
