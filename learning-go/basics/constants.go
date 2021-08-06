package main

import (
	"fmt"
	"math"
)

const s string = "Nachiket Kanore"

func main() {
	fmt.Println(s)

	const n = 100002211

	const d = 1e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}
