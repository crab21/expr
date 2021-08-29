package ast

import (
	"regexp"
	"strings"
)

func ParseBrackets(exp string) {
	rex := regexp.MustCompile(`\(([^)]+)\)`)
	out := rex.FindAllStringSubmatch(exp, -1)
	for _, i := range out {

		exp = strings.ReplaceAll(exp, "("+i[1]+")", "2==2")
	}

}
