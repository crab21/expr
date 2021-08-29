package ast

import (
	"errors"
	"os"
	"strings"
)

// 生成一个 AST 结构指针
func NewAST(toks []*Token, s string) *AST {
	a := &AST{
		Tokens: toks,
		source: s,
	}
	if a.Tokens == nil || len(a.Tokens) == 0 {
		a.Err = errors.New("empty token")
	} else {
		a.currIndex = 0
		a.currTok = a.Tokens[0]
	}
	return a
}

func Exec(exp string) bool {
	// exp := "1==1 && gogoowang==gogoowang"

	// input text -> []token
	toks, _ := Parse(exp)
	//if err != nil {
	//	//fmt.Println("ERROR: " + err.Error())
	//	return
	//}
	// []token -> AST Tree
	ast := NewAST(toks, exp)
	if ast.Err != nil {

		return false
	}
	// AST builder
	ar := ast.ParseExpression()
	if ast.Err != nil {

		return false
	}
	// fmt.Printf("ExprAST: %+v\n", ar)

	// AST traversal -> result
	r, paths := ExprASTResult(ar, nil)

	for _, v := range paths {
		os.Remove(v)
	}
	if v, ok := r.(string); ok {
		if strings.TrimSpace(v) == "" || strings.TrimSpace(v) == "false" {
			return false
		}
		if strings.TrimSpace(v) == "true" {
			return true
		}
	}
	return false
}

func SaveResultToTektonPath(tektonPath string, r bool) {
	f, err := os.OpenFile(tektonPath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("open file error" + err.Error())
	}
	defer f.Close()
	//rr := strings.ReplaceAll(r.(string),"\n","")
	if r {
		f.WriteString("true")
	} else {
		f.WriteString("false")
	}

}
