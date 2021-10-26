package main

import "fmt"

func main() {
	var n = 9
	var ip *int

	ip = &n

	fmt.Printf("address of n is %d\n" , ip)
	fmt.Printf("address of n is %d\n" , &n)
	fmt.Printf("value of n is %d\n" , *ip)
}
