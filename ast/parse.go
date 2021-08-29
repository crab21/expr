package ast

import (
	"errors"
	"fmt"
	"strings"
)

func (p *Parser) parse() []*Token {
	toks := make([]*Token, 0)
	for {
		tok := p.nextTok()
		if tok == nil {
			break
		}
		//////fmt.Println("tok:======>", tok.Tok, "type:====>", tok.Type, "offset:=====>", tok.Offset)
		toks = append(toks, tok)
	}

	//////fmt.Println("=======================parse end=======================")
	return toks
}

func (p *Parser) nextTok() *Token {
	if p.Offset >= len(p.Source) || p.Err != nil {
		return nil
	}

	var err error
	// 跳过所有无意义的空白符
	for p.isWhitespace(p.Ch) && err == nil {
		err = p.nextCh()
	}

	start := p.Offset
	var tok *Token
	switch p.Ch {
	case
		'(',
		')':
		tok = &Token{
			Tok:  p.Source[start : p.Offset+1],
			Type: Operator,
		}
		tok.Offset = start
		p.Err = p.nextCh()
	case
		// '+',
		// '-',
		// '*',
		// '/',
		// '^',
		// '%',
		'=',
		'>',
		'!',
		'&',
		'|',
		'<':
		for {
			a, n := p.OperatorContinue(p.Ch)
			if a > 0 && n {
				if p.nextCh() == nil {
					break
				}
				if a == 1 && n {
					p.nextCh()
					continue
				}
			}
			if a == -2 {
				if n {
					p.nextCh()
				}
				break
			}
			if a == -1 {

				if n {
					c := p.Source[p.Offset]

					if p.isDigitNum(c) {
						p.beforeCh()
						break
					}
					p.nextCh()
					break
				} else {
					p.nextCh()
					break
				}

			}
			if a == -3 {
				break
			}
		}
		tok = &Token{
			Tok:  p.Source[start : p.Offset+1],
			Type: Operator,
		}
		tok.Offset = start
		p.Err = p.nextCh()
	case
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'$', '\'', '{', '}', '"', '#':
		for {
			// result := p.isDigitNum(p.Ch) && p.nextCh() == nil
			//////fmt.Println(string(p.Ch), string(p.GetAfterCh()), GetTokPrecedence(string(p.GetAfterCh())))
			//if p.GetAfterCh() == 0 || p.GetAfterCh() == ' ' || GetTokPrecedence(string(p.GetAfterCh())) {
			if p.GetAfterCh() == 0 || GetTokPrecedence(string(p.GetAfterCh())) {
				p.nextCh()
				break
			}
			p.nextCh()
		}
		tok = &Token{
			Tok:  strings.ReplaceAll(p.Source[start:p.Offset], "_", ""),
			Type: Literal,
		}
		tok.Offset = start

		// 捕获错误
	default:
		if p.Ch != ' ' {
			s := fmt.Sprintf("symbol error: unkown '%v', pos [%v:]\n%s",
				string(p.Ch),
				start,
				ErrPos(p.Source, start))
			p.Err = errors.New(s)
		}
	}
	return tok
}

// 前进到下一个字符
func (p *Parser) nextCh() error {
	p.Offset++

	if p.Offset < len(p.Source) {
		p.Ch = p.Source[p.Offset]
		return nil
	}

	// 到达字符串末尾
	return errors.New("after EOF now ch:" + string(p.Ch))
}

// 前进到下一个字符
func (p *Parser) beforeCh() error {
	p.Offset--

	if p.Offset < len(p.Source) {
		p.Ch = p.Source[p.Offset]
		return nil
	}

	// 到达字符串末尾
	return errors.New("before EOF")
}

func (p *Parser) GetbeforeCh() byte {
	cur := p.Offset
	if cur == 0 {
		return 0
	}
	cur--
	return p.Source[cur]

}

func (p *Parser) GetAfterCh() byte {
	cur := p.Offset
	cur++
	if len(p.Source) <= cur {
		return 0
	}
	return p.Source[cur]

}

// 空白符
func (p *Parser) isWhitespace(c byte) bool {
	return c == ' ' ||
		c == '\t' ||
		c == '\n' ||
		c == '\v' ||
		c == '\f' ||
		c == '\r'
}

// 数字
func (p *Parser) isDigitNum(c byte) bool {
	return '0' <= c && c <= '9' || c == '.' || c == '_' || c == 'e'
}

// 数字
func (p *Parser) OperatorContinue(c byte) (int, bool) {
	if c == '!' {
		after := p.GetAfterCh()
		if after == 0 {
			return 1, true
		}
		if after == '=' {
			return -1, true
		}
		return -2, false
	}
	if c == '<' {
		after := p.GetAfterCh()
		if after == 0 {
			return 1, true
		}
		if after == '=' {
			return -1, true
		}
		return -2, false
	}

	if c == '>' {
		after := p.GetAfterCh()

		if after == 0 {
			return 1, true
		}
		if after == '=' {
			return -1, true
		}
		return -2, false
	}

	if c == '&' {
		before := p.GetbeforeCh()
		if before == 0 {
			return 1, false
		}
		if before == byte(' ') {
			return -1, false
		}
		if before == '&' {
			return 1, true
		}
		return 1, true
	}

	if c == '|' {
		before := p.GetbeforeCh()
		if before == 0 {
			return 1, false
		}
		if before == byte(' ') {
			return -1, false
		}
		if before == '|' {
			return -2, false
		}
		return 1, true
	}

	if c == '=' {
		before := p.GetbeforeCh()
		if before == 0 {
			return 1, false
		}

		if before == '=' || before == '<' || before == '>' || before == '!' {
			return -1, true
		}
		return -2, true
	}
	return -3, true
}

// 对错误包装，进行可视化展示
func ErrPos(s string, pos int) string {
	r := strings.Repeat("-", len(s)) + "\n"
	s += "\n"
	for i := 0; i < pos; i++ {
		s += " "
	}
	s += "^\n"
	return r + s + r
}

// 封装词法分析过程，直接调用该函数即可解析字符串为[]Token
func Parse(s string) ([]*Token, error) {
	// 初始化 Parser
	//fmt.Println(s, "===============>ssssssssssssss")
	p := &Parser{
		Source: s,
		Err:    nil,
		Ch:     s[0],
	}
	// 调用 parse 方法
	toks := p.parse()
	//if p.Err != nil {
	//	//fmt.Println("====================end=============")
	//	return nil, p.Err
	//}
	return toks, nil
}

/////////////////////////////////////

// 定义操作符优先级，value 越高，优先级越高
var precedence = map[string]int{
	// "+": 20, "-": 20,
	// "*": 40, "/": 40, "%": 40,
	// "^":  60,
	"==": 70, "<=": 70, ">=": 70, "!=": 70,
	"<": 70, ">": 70,
	"&&": 69, "||": 69,
	"(": 80, ")": 80, "!": 90,
}

var precedenceMap = map[string]interface{}{
	// "+": 20, "-": 20,
	// "*": 40, "/": 40, "%": 40,
	// "^":  60,
	"==": 70, "<=": 70, ">=": 70, "!=": 70,
	"<": 70, ">": 70,
	"&&": 69, "||": 69,
	"(": 80, ")": 80, "!": 90,
	"=": 0, "&": 0, "|": 0,
}

func GetTokPrecedence(k string) bool {
	if _, ok := precedenceMap[k]; ok {
		return true
	}
	return false
}

// 语法分析器入口
func (a *AST) ParseExpression() ExprAST {
	lhs := a.parsePrimary()
	return a.parseBinOpRHS(0, lhs)
}

// 获取下一个 Token
func (a *AST) getNextToken() *Token {
	a.currIndex++
	if a.currIndex < len(a.Tokens) {
		a.currTok = a.Tokens[a.currIndex]
		return a.currTok
	}
	return nil
}

// 获取操作优先级
func (a *AST) getTokPrecedence() int {
	if p, ok := precedence[a.currTok.Tok]; ok {
		return p
	}
	return -1
}

// 解析数字，并生成一个 NumberExprAST 节点
func (a *AST) parseNumber() NumberExprAST {
	// f64, err := strconv.ParseFloat(a.currTok.Tok, 64)
	// if err != nil {
	// 	//fmt.Println(a.currTok.Tok, "==================>")
	// 	a.Err = errors.New(
	// 		fmt.Sprintf("%v\nwant '(' or '0-9' but  offset:%d get %s %s",
	// 			err.Error(),
	// 			a.currTok.Offset,
	// 			a.currTok.Tok,
	// 			ErrPos(a.source, a.currTok.Offset)))
	// 	return NumberExprAST{}
	// }
	n := NumberExprAST{
		Val: a.currTok.Tok,
	}
	a.getNextToken()
	return n
}

// 获取一个节点，返回 ExprAST
// 这里会处理所有可能出现的类型，并对相应类型做解析
func (a *AST) parsePrimary() ExprAST {
	switch a.currTok.Type {
	case Literal:
		return a.parseNumber()
	case Operator:
		// 对 () 语法处理
		//if a.currTok.Tok == "(" {
		_ = a.getNextToken()
		//fmt.Println("tttt========>", tt)
		e := a.ParseExpression()
		if e == nil {
			return nil
		}
		//if a.currTok.Tok != ")" {
		//	a.Err = errors.New(
		//		fmt.Sprintf("want ')' but get %+v\n%s",
		//			a.currTok,
		//			ErrPos(a.source, a.currTok.Offset)))
		//	return nil
		//}
		a.getNextToken()
		return e
		//} else {
		//	//fmt.Println("else.......",a.currTok,a.currTok.Tok,a.currTok.Offset,a.currTok.Type)
		//	//e := a.ParseExpression()
		//	a.getNextToken()
		//	return nil
		//}
	default:
		return nil
	}
}

// 循环获取操作符的优先级，将高优先级的递归成较深的节点
// 这是生成正确的 AST 结构最重要的一个算法，一定要仔细阅读、理解
func (a *AST) parseBinOpRHS(execPrec int, lhs ExprAST) ExprAST {
	for {
		tokPrec := a.getTokPrecedence()
		if tokPrec < execPrec {
			return lhs
		}
		binOp := a.currTok.Tok
		if a.getNextToken() == nil {
			return lhs
		}
		rhs := a.parsePrimary()
		if rhs == nil {

			lhs = BinaryExprAST{
				Op:  binOp,
				Lhs: lhs,
				Rhs: rhs,
			}
			return lhs
		}
		nextPrec := a.getTokPrecedence()
		if tokPrec < nextPrec {
			// 递归，将当前优先级+1
			rhs = a.parseBinOpRHS(tokPrec+1, rhs)
			if rhs == nil {
				return nil
			}
		}
		lhs = BinaryExprAST{
			Op:  binOp,
			Lhs: lhs,
			Rhs: rhs,
		}
	}
}
