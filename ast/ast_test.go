package ast

import (
	"testing"
)

func TestExec(t *testing.T) {
	// exp := "1==1 && gogoowang==gogoowang"

	// v:=Exec(exp)
	//////fmt.Println(v)
	// // for i := 0; i < 26; i++ {
	// // 	a := 97 + i
	// // 	fmt.Print("'", string(rune(a)), "',")
	// // }
	a := "a"
	b := "b"
	if a > b {
		t.Errorf("ok")
	} else {
		t.Error("no")
	}
}
