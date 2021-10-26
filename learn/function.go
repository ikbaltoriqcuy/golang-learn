package main

import "fmt"

func main() {
	fmt.Printf(" %d \n" , function1(3,4))
	fmt.Printf(" %d \n" , function2(3,4))

	var a,b = function3(3,"ikbal")
	fmt.Printf(" %s %d \n" , a, b)
}


func function1(n int, s int) int {
	return n+s
}

func function2(n , s int) int {
	return n+s
}


func function3(n int , s string) (string, int) {
	return s,n
}
