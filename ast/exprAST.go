package ast

import (
	"io/ioutil"
	"os"
	"time"

	"cal"
)

func ExprASTResult(expr ExprAST, pps []string) (lll interface{}, paths []string) {
	var l, r interface{}

	switch expr.(type) {
	case BinaryExprAST:
		ast := expr.(BinaryExprAST)
		// 递归左节点
		ls, pathsl := ExprASTResult(ast.Lhs, paths)
		l = ls
		paths = append(paths, pathsl...)
		// 递归右节点
		rs, pathsr := ExprASTResult(ast.Rhs, paths)
		r = rs
		paths = append(paths, pathsr...)
		// if l == r && (ast.Op == "&&" || ast.Op == "||") {
		// 	return true
		// }
		// if l == r && (ast.Op == "&&" || ast.Op == "||") {
		// 	return true
		// }
		time.Sleep(1 * time.Millisecond)
		swd, _ := os.Getwd()
		result := GetUID()
		pathName := swd + "/" + result

		a := cal.CompareValueAndOperator{
			UniqueID:      pathName,
			Values:        []interface{}{r.(string)},
			CompareSymbol: "all",
			BeCompareValues: []cal.CompareOperator{
				{
					CompareValue:      l.(string),
					Operator:          ast.Op,
					TwoValuesOperator: "",
				},
			},
		}

		cal.CalByStruct(a)

		// sv := swd + "/expr" + " -uniqueID \"" + pathName + "\" -compareValue \"" +
		// l.(string) + "\" -operator \"" + ast.Op + "\" -values \"" + "[\\\"" + r.(string) + "\\\"]\""

		// cmd := exec.Command("/bin/bash", "-c", sv)
		// var out bytes.Buffer
		// var stderr bytes.Buffer
		// cmd.Stdout = &out
		// cmd.Stderr = &stderr
		// _ = cmd.Run()

		f, _ := os.Open(pathName)
		defer f.Close()
		ress, _ := ioutil.ReadAll(f)

		paths = append(paths, pathName)

		return string(ress), paths
	// 现在 l,r 都有具体的值了，可以根据运算符运算
	// switch ast.Op {
	// case "+":
	// 	return l + r
	// case "-":
	// 	return l - r
	// case "*":
	// 	return l * r
	// case "/":
	// 	if r == 0 {
	// 		panic(errors.New(
	// 			fmt.Sprintf("violation of arithmetic specification: a division by zero in ExprASTResult: [%g/%g]",
	// 				l,
	// 				r)))
	// 	}
	// 	return l / r
	// case "%":
	// 	return float64(int(l) % int(r))
	// default:
	// }
	case NumberExprAST:

		return expr.(NumberExprAST).Val, paths
	}

	return 0.0, paths

}
