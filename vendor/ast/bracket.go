package ast

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseBrackets(exp string) {
	rex := regexp.MustCompile(`\(([^)]+)\)`)
	out := rex.FindAllStringSubmatch(exp, -1)
	for _, i := range out {
		fmt.Println(i[1])
		exp = strings.ReplaceAll(exp, "("+i[1]+")", "2==2")
	}

	fmt.Println("exp", exp)

}
