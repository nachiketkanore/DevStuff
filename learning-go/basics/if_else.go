package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Println("even", i)
		}
		if i % 2 == 1 {
			fmt.Println("odd", i)
		}
	}

	if true && true && true && true {
		fmt.Println("True")
	}
}
