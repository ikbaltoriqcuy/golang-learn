package main

import (
	"encoding/json"
	"fmt"
)

type Bioku struct{
	Id int
	Name string
	Age int
}

func main() {
	n := &Bioku{Id:32,Name:"ikbal",Age:21}
	b,er := json.Marshal(n)
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println(string(b) + "\n")

	var k Bioku
	k.Age = 23
	k.Name = "lsd"
	k.Id = 12

	s,er := json.Marshal(k)
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println(string(s))
}
