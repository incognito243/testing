package main

import "fmt"

func main() {
	a, b, c := 100.0, 4.0, 5.0
	fmt.Printf(checkTriangleType(a, b, c))
}

func checkTriangleType(a, b, c float64) string {
	if a+b > c && a+c > b && b+c > a {
		if a == b && b == c {
			return "Tam giác đều"
		} else if a == b || b == c || c == a {
			return "Tam giác cân"
		} else {
			return "Tam giác thường"
		}
	} else {
		return "Không phải tam giác"
	}
}
