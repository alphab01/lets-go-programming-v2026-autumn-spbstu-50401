package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var s string
	_, err = fmt.Scan(&s)
	if err != nil {
		fmt.Println("Error, cant read string")
    return
	}
	if s == "+" {
		fmt.Println(a + b)
	} else if s == "-" {
		fmt.Println(a - b)
	} else if s == "/" {
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("Division by zer  o")
		}
	} else if s == "*" {
		fmt.Println(a * b)
	} else {
		fmt.Println("Invalid operation")
	}
}
