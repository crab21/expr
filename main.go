package main

import (
	"ast"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var exprValue string

const pathEnd = "/"

type Exprvalue struct {
	ValueCompare string
	TektonPath   string
	NopFlag      bool
}

type ResultInfo struct {
	ResultJobID string
	Result      bool
}

func main() {
	// exprValue = "W3siVmFsdWVDb21wYXJlIjoiZXlKbGVIQnlWbUZzZFdWeklqcGJJakU5UFRFaVhYMD0iLCJUZWt0b25QYXRoIjoiL3Rla3Rvbi9yZXN1bHRzLzRkMGEtOTU4MSIsIk5vcEZsYWciOnRydWV9XQ=="

	flag.StringVar(&exprValue, "exprValue", "", "eg: base64 ")
	flag.Parse()

	value, err := base64.StdEncoding.DecodeString(exprValue)
	if err != nil {
		panic("parse base64 exprValue" + err.Error())
	}

	ev := make([]Exprvalue, 0)
	_ = json.Unmarshal([]byte(value), &ev)
	resultPath := make([]ResultInfo, 0, len(ev))
	for _, v := range ev {
		var chooseValuePath string

		if !v.NopFlag {
			chooseValuePath = os.Getenv("DELIVER")
		} else {
			chooseValuePath = "#{option}"
		}
		resultPath = evalValue(resultPath, v.ValueCompare, v.TektonPath, chooseValuePath)
	}

	vResult, _ := json.Marshal(resultPath)
	fmt.Println(string(vResult))
}

func evalValue(resultPath []ResultInfo, expr string, tektonPath string, chooseValuePath string) []ResultInfo {

	if expr == "" || tektonPath == "" || chooseValuePath == "" {
		panic("expr/tektonPath/chooseValuePath is not allow empty")
	}
	exprvalue, err := base64.StdEncoding.DecodeString(expr)
	if err != nil {
		panic("parse base64 exprValue" + err.Error())
	}

	mp := make(map[string][]string)
	err = json.Unmarshal(exprvalue, &mp)
	if err != nil {
		panic("Unmarshal  exprValue" + err.Error())
	}

	var optionExpr string = "#{option}"
	if v, ok := mp["optionExpr"]; ok {
		optionExpr = v[0]
	}

	result := chooseValuePath

	exprvalues := mp["exprValues"]
	var resultFalg bool
	for _, v := range exprvalues {
		resultv := strings.ReplaceAll(v, optionExpr, string(result))
		resultv = strings.ReplaceAll(resultv, " ", "")
		if resultv == "" {
			continue
		}
		rv := ast.Exec(resultv)
		//////fmt.Println("rv for value:----------->", rv, "resultFalg--->", resultFalg)
		resultFalg = resultFalg || rv
	}

	//fmt.Println("rv==============>", resultFalg)
	lastIndex := strings.LastIndex(tektonPath, "/")
	if lastIndex != -1 {
		resultPath = append(resultPath, ResultInfo{strings.TrimRight(tektonPath[lastIndex+1:], "-expression-step"), resultFalg})
	}

	ast.SaveResultToTektonPath(tektonPath, resultFalg)
	return resultPath
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
