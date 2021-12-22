package ast

import (
	"fmt"
	"testing"
	"unicode/utf16"
)

func TestWords(t *testing.T) {

	s := utf16.Encode([]rune("王"))
	fmt.Sprintf("%v", s)
}
