package main

import "fmt"

type bio struct{
	id int
	name string
	age int
}


func main() {
	var n bio
	n.id = 34
	n.name = "ikbal"
	n.age = 24

	fmt.Printf("id : %d " , n.id)
	fmt.Printf("name : %s " , n.name)
	fmt.Printf("age : %d " , n.age)
}
