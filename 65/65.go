package main

import (
	"fmt"
	"regexp"
)

func isNumber(s string) bool {
	regexpStr:=`^ *[+-]?(\d+|\d+\.\d+|\d+\.|\.\d+)(e[+-]?\d+)? *$`
	match,_:=regexp.MatchString(regexpStr,s)
	return match
}

func main() {
	fmt.Println(isNumber(".3e2"))

}
