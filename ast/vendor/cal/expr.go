package cal

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sony/sonyflake"
)

type Tweet struct {
	Len string
}

type Env struct {
	Tweets []Tweet
}

type CompareValueAndOperator struct {
	UniqueID        string
	Values          []interface{}
	BeCompareValues []CompareOperator
	CompareSymbol   string
}

type CompareOperator struct {
	CompareValue      interface{}
	Operator          interface{}
	TwoValuesOperator interface{}
}

func main() {
	// a := CompareValueAndOperator{
	// 	UniqueID:      "1627662612",
	// 	Values:        []interface{}{"gogo"},
	// 	CompareSymbol: "all",
	// 	BeCompareValues: []CompareOperator{
	// 		{
	// 			CompareValue:      "gogo",
	// 			Operator:          "==",
	// 			TwoValuesOperator: "&&",
	// 		},
	// 	},
	// }
	// fmt.Println(CompareValues((a)))
	// sText := "&&"
	// textQuoted := strconv.QuoteToASCII(sText)
	// textUnquoted := textQuoted[1 : len(textQuoted)-1]
	// fmt.Println(textUnquoted)
	// ss, _ := os.Getwd()
	// fmt.Println(ss)
	// SaveToResult(true, "task1-output1")
	var values, compareValue, uniqueID, operator, TwoValuesOperator, CompareSymbol string
	flag.StringVar(&values, "values", "", `eg: ["a","b"]`)
	flag.StringVar(&compareValue, "compareValue", "", `"a"`)
	flag.StringVar(&uniqueID, "uniqueID", "", ``)
	flag.StringVar(&operator, "operator", "", ``)
	flag.StringVar(&TwoValuesOperator, "TwoValuesOperator", "", `eg: &&/||`)
	flag.StringVar(&CompareSymbol, "CompareSymbol", "all", `eg: one/all/any`)
	flag.Parse()

	if values == "" || compareValue == "" || operator == "" {
		fmt.Println("empty value")
		panic("empty value")
	}

	fmt.Println(values)

}

func CalResult(values string, uniqueID, CompareSymbol, compareValue, operator,
	TwoValuesOperator string) {
	fmt.Println("enter...........")
	value := make([]string, 0)
	_ = json.Unmarshal([]byte(values), &value)
	fmt.Println(value)

	a := CompareValueAndOperator{
		UniqueID: uniqueID,

		CompareSymbol: CompareSymbol,
		BeCompareValues: []CompareOperator{
			{
				CompareValue:      compareValue,
				Operator:          operator,
				TwoValuesOperator: TwoValuesOperator,
			},
		},
	}
	for _, v := range value {
		a.Values = append(a.Values, v)
	}
	fmt.Println(a)
	uinque, result, _ := CompareValues((a))

	SaveToResult(result, uinque)
}

func CalByStruct(a CompareValueAndOperator) {
	uinque, result, _ := CompareValues((a))

	SaveToResult(result, uinque)
}

func CompareValues(a CompareValueAndOperator) (unique string, result bool, errRet error) {
	var (
		uniqueID string = a.UniqueID

		v = a.BeCompareValues[0]
	)

	switch v.Operator.(string) {
	case "&&", "&", "||", "|":
		var values []bool = make([]bool, 0, len(a.Values))
		for _, v := range a.Values {
			cc := v.(string)
			if cc == "true" || cc == "false" {
				vs, _ := strconv.ParseBool(cc)
				values = append(values, vs)
			}
		}
		var vb bool
		cvv := v.CompareValue.(string)
		if cvv == "true" || cvv == "false" {
			vs, _ := strconv.ParseBool(cvv)
			vb = vs
		}
		result = CompareBool(values, vb, v.Operator.(string), v.TwoValuesOperator.(string), a.CompareSymbol)
		fmt.Println("compare value: ", vb, "  v  values:", values, "operator:", v.Operator.(string))
	default:
		var values []string = make([]string, 0, len(a.Values))
		for _, v := range a.Values {
			var sc string
			sc = strings.TrimLeft(v.(string), "'")
			sc = strings.TrimSpace(sc)
			sc = strings.TrimLeft(sc, "\"")
			sc = strings.TrimLeft(sc, "\n")
			sc = strings.TrimRight(sc, "'")
			sc = strings.TrimRight(sc, "\"")
			sc = strings.TrimRight(sc, "\n")
			values = append(values, sc)
		}

		var sc string
		sc = strings.TrimLeft(v.CompareValue.(string), "'")
		sc = strings.TrimSpace(sc)
		sc = strings.TrimLeft(sc, "\"")
		sc = strings.TrimLeft(sc, "\n")
		sc = strings.TrimRight(sc, "'")
		sc = strings.TrimRight(sc, "\"")
		sc = strings.TrimRight(sc, "\n")
		fmt.Println("compare value: ", sc, "  v  values:", values, "operator:", v.Operator.(string))
		result = CompareString(values, sc, v.Operator.(string), v.TwoValuesOperator.(string), a.CompareSymbol)
	}

	if uniqueID == "" {
		n, _ := sonyflake.NewSonyflake(sonyflake.Settings{}).NextID()
		result := strconv.FormatUint(n, 10)
		uniqueID = result
	}

	return uniqueID, result, nil
}

func SaveToResult(result bool, pathName string) {
	f, err := os.OpenFile(pathName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	defer f.Close()
	if result {
		f.Write([]byte("true"))
	} else {
		f.Write([]byte("false"))
	}

}
