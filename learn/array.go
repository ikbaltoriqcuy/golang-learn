package main

import (
	"fmt"
	"strconv"
)

var g [100]string

func main() {
	for i:=0; i<100 ; i++ {
		g[i] = "number " + strconv.Itoa(i)
		fmt.Printf("%s\n", g[i])
	}
}
