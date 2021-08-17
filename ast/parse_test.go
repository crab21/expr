package ast

import (
	"testing"
)

func TestParse(t *testing.T) {
	s := "1==2 && (2==3)"
	Parse(s)
}
