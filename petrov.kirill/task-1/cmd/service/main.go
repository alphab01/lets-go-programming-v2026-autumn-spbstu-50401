package main

import "fmt"

func main() {
  var a, b int
  _, err := fmt.Scan(&a)
  if (err != nil) {
    fmt.Println("Invalid first operand")
    return
  }
  _, err = fmt.Scan(&b)
  if (err != nil) {
    fmt.Println("Invalid second operand")
    return
  }
  var s string
  fmt.Scan(&s)
  if (s == "+") {
    fmt.Println(a + b)
  } else if (s == "-") {
    fmt.Println(a - b)
  } else if (s == "/") {
    fmt.Println(a / b)
  } else {
    fmt.Println(a * b)
  }
}
