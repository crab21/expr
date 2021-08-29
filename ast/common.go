package ast

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

const (
	// two types
	Literal = iota
	Operator
)

type Token struct {
	// origin string
	Tok string

	// Literal or Opertor
	Type   int
	Offset int
}

type Parser struct {
	// input or source string
	Source string

	// current scan byte
	Ch byte

	// current scan offset
	Offset int

	Err error
}

// 基础表达式节点接口
type ExprAST interface {
	toStr() string
}

// 数字表达式节点
type NumberExprAST struct {
	// 具体的值
	Val interface{}
}

// 操作表达式节点
type BinaryExprAST struct {
	// 操作符
	Op string
	// 左右节点，可能是 数字表达式节点/操作表达式节点/nil
	Lhs, Rhs ExprAST
}

// 实现接口
func (n NumberExprAST) toStr() string {
	return fmt.Sprintf(
		"NumberExprAST:%s",
		// strconv.FormatFloat(n.Val, 'f', 0, 64),
		n.Val,
	)
}

// 实现接口
func (b BinaryExprAST) toStr() string {
	return fmt.Sprintf(
		"BinaryExprAST: (%s %s %s)",
		b.Op,
		b.Lhs.toStr(),
		b.Rhs.toStr(),
	)
}

// AST 生成器结构体
type AST struct {
	// 词法分析的结果
	Tokens []*Token
	// 源字符串
	source string
	// 当前分析器分析的 Token
	currTok *Token
	// 当前分析器的位置
	currIndex int
	// 错误收集
	Err error
}

func GetUID() string {
	node, err := snowflake.NewNode(10)
	if err != nil {

		return ""
	}
	id := node.Generate().Int64()
	return strconv.FormatInt(id, 10)

}
