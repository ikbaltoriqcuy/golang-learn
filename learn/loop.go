package main

import "fmt"

func main() {
	for i:=1; i<10; i++ {
		fmt.Printf("%d", i)
	}

	var n = 9
	for true {
		n++
		fmt.Printf("%d", n)
		if (n > 15) {
			break
		}
	}
}
