package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = "ikbal toriq"
	fmt.Printf("Length of string \"%s\" is %d\n", data, len(data))

	var n = "ikbal"
	var s = "toriq"
	fmt.Printf("It's same of string \"%s\" and \"%s\" is %v\n", n, s , strings.Compare(n ,s))

	//split

	var text = "hello wonderfull world"
	var b= []rune(text)
	fmt.Printf("%s", string(b[5:10]))
}
