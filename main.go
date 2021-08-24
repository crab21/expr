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

func main() {

	// valueCompare := map[string][]string{
	// 	"exprValues": []string{"#{option}<b"},
	// }
	// marshal, _ := json.Marshal(valueCompare)
	// toString := base64.StdEncoding.EncodeToString(marshal)
	// exprvalue, err := base64.StdEncoding.DecodeString(toString)
	// tektonPathTmp, _ := os.Getwd()
	// tektonPath = tektonPathTmp + "gogo-test"
	// chooseValuePath = "a"
	// exprValue = "W3siVmFsdWVDb21wYXJlIjoiZXlKbGVIQnlWbUZzZFdWeklqcGJJakU5UFRFaVhYMD0iLCJUZWt0b25QYXRoIjoiL3Rla3Rvbi9yZXN1bHRzLzQxMTctYmE4NC1leHByZXNzaW9uLXN0ZXAiLCJOb3BGbGFnIjp0cnVlfSx7IlZhbHVlQ29tcGFyZSI6ImV5SmxlSEJ5Vm1Gc2RXVnpJanBiSWpFaFBURWlYWDA9IiwiVGVrdG9uUGF0aCI6Ii90ZWt0b24vcmVzdWx0cy80YjM2LWJhZmYtZXhwcmVzc2lvbi1zdGVwIiwiTm9wRmxhZyI6dHJ1ZX1d"

	flag.StringVar(&exprValue, "exprValue", "", "eg: base64 ")
	flag.Parse()

	value, err := base64.StdEncoding.DecodeString(exprValue)
	if err != nil {
		panic("parse base64 exprValue" + err.Error())
	}

	ev := make([]Exprvalue, 0)
	_ = json.Unmarshal([]byte(value), &ev)
	for _, v := range ev {
		var chooseValuePath string
		if !v.NopFlag {
			chooseValuePath = os.Getenv("DELIVER")
		} else {
			chooseValuePath = "#{option}"
		}
		evalValue(v.ValueCompare, v.TektonPath, chooseValuePath)
	}
}

func evalValue(expr string, tektonPath string, chooseValuePath string) {

	if expr == "" || tektonPath == "" || chooseValuePath == "" {
		panic("expr/tektonPath/chooseValuePath is not allow empty")
	}
	exprvalue, err := base64.StdEncoding.DecodeString(expr)
	if err != nil {
		panic("parse base64 exprValue" + err.Error())
	}

	fmt.Println(exprvalue)
	mp := make(map[string][]string)
	err = json.Unmarshal(exprvalue, &mp)
	if err != nil {
		panic("Unmarshal  exprValue" + err.Error())
	}
	fmt.Println(err, mp)

	var optionExpr string = "#{option}"
	if v, ok := mp["optionExpr"]; ok {
		optionExpr = v[0]
	}

	result := chooseValuePath

	exprvalues := mp["exprValues"]
	var resultFalg bool
	fmt.Println("default value:-->", resultFalg)
	for _, v := range exprvalues {
		resultv := strings.ReplaceAll(v, optionExpr, string(result))
		fmt.Println("exprValue: ", resultv)
		resultv = strings.ReplaceAll(resultv, " ", "")
		if resultv == "" {
			continue
		}
		rv := ast.Exec(resultv)
		// fmt.Println("rv for value:----------->", rv, "resultFalg--->", resultFalg)
		resultFalg = resultFalg || rv
	}

	fmt.Println("rv==============>", resultFalg)

	ast.SaveResultToTektonPath(tektonPath, resultFalg)

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
