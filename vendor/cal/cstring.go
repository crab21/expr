package cal

import (
	"fmt"

	"github.com/antonmedv/expr"
)

type TweetString struct {
	Len string
}

type EnvString struct {
	Tweets []TweetString
}

func smain() {
	CompareString([]string{"gogo"}, "gogo", "==", "", "all")
}

func CompareString(value []string, compareValue string, operator string, TwoValuesOperator string,
	CompareSymbol string) bool {

	if v, ok := TurnOperatorToCompare[operator]; ok {
		operator = v
	}

	s := fmt.Sprintf("all(Tweets, {.Len %s \"%s\"})", operator, compareValue)
	code := fmt.Sprintf(`%s`, s)

	program, err := expr.Compile(code, expr.Env(EnvString{}))
	if err != nil {
		panic(err)
	}

	tws := make([]TweetString, 0, len(value))
	for _, v := range value {
		tws = append(tws, TweetString{v})
	}
	env := EnvString{
		Tweets: tws,
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	//fmt.Println(output)
	return output.(bool)
}

var TurnOperatorToCompare = map[string]string{
	">":  "<",
	">=": "<=",
	"<":  ">",
	"<=": ">=",
}
