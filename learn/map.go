package main

import "fmt"

func main() {
	var data map[string]int
	data = make(map[string]int)

	data["ikbal"] =  4
	data["toriq"] = 5
	data["saya"] = 5

	fmt.Printf("value of data[\"ikbal\"] is %d"  , data["ikbal"])
}
