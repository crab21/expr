package cal

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCompareValues(t *testing.T) {

	a := CompareValueAndOperator{
		UniqueID:      "1627662612",
		Values:        []interface{}{"true"},
		CompareSymbol: "all",
		BeCompareValues: []CompareOperator{
			{
				CompareValue:      true,
				Operator:          "and",
				TwoValuesOperator: "&&",
			},
		},
	}
	v, _ := json.Marshal(a)
	fmt.Println(CompareValues((a)))
	fmt.Println(string(v))

	b := `{
		"UniqueID": "1627662612",
		"Values": [
		 "gogo"
		],
		"BeCompareValues": [
		 {
		  "CompareValue": "gogo",
		  "Operator": "==",
		  "TwoValuesOperator": "&&"
		 }
		],
		"CompareSymbol": "all"
	   }`

	bs := &CompareValueAndOperator{}
	_ = json.Unmarshal([]byte(b), bs)
	fmt.Println(bs.BeCompareValues[0].TwoValuesOperator)
}
