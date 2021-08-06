package main

import "fmt"

func main() {

	i := 1
	fmt.Println("initial value ... i =", i)

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 20; j <= 30; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Infinite loop")
		break
	}

	for n := 0; n <= 20; n++ {
		if n % 2 == 1 {
			fmt.Println("odd number", n)
		}
	}

}
