package cal

type TweetBool struct {
	Len bool
}

type EnvBool struct {
	Tweets []TweetBool
}

func ssmain() {
	CompareBool([]bool{true}, true, "&&", "", "all")
}

// func CompareBool(value []bool, compareValue bool, operator string, TwoValuesOperator string,
// 	CompareSymbol string) bool {
// 	s := fmt.Sprintf("all(Tweets, {.Len && %v})", compareValue)
// 	code := fmt.Sprintf(`%s`, s)

// 	program, err := expr.Compile(code, expr.Env(EnvBool{}))
// 	if err != nil {
// 		panic(err)
// 	}

// 	tws := make([]TweetBool, 0, len(value))
// 	for _, v := range value {
// 		tws = append(tws, TweetBool{v})
// 	}
// 	env := EnvBool{
// 		Tweets: tws,
// 	}

// 	output, err := expr.Run(program, env)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//fmt.Println(output)
// 	return output.(bool)
// }
func CompareBool(value []bool, compareValue bool, operator string, TwoValuesOperator string,
	CompareSymbol string) (result bool) {

	switch operator {
	case "&&":
		result = compareValue && value[0]
	case "||":
		result = compareValue || value[0]

	default:
		result = false
	}
	//fmt.Println(result)
	return
}
