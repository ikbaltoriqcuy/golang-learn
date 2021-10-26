package main

import "fmt"

var n = 34
var k = 3

func main() {
	n += 3
	fmt.Printf("data = %d\n", n)

	n -= 1
	fmt.Printf("data = %d\n", n)

	n *= 2
	fmt.Printf("data = %d\n", n)

	n /= 2
	fmt.Printf("data = %d\n", n)

	fmt.Printf("(%d = %d)= %v\n", n , k , n == k)

	fmt.Printf("(%d < %d)= %v\n",  n , k , n < k)

	fmt.Printf("(%d > %d)= %v\n",  n , k , n > k)

	fmt.Printf("(%d != %d)= %v\n",  n , k , n != k)
}