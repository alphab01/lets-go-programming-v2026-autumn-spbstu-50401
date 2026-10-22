package main

import (
  "fmt"
)

func main() {
  var a, b int
  fmt.Scan(&a)
  var d []int
  for i := 0; i < a; i++ {
    fmt.Scan(&b)
    d = append(d, b)
  }
}
