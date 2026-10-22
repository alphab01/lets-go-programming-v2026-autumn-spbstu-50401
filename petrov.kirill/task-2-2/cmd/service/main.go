package main

import (
  "fmt"
)

func gsort(d []int) []int {
  i := 0
  for ; i < len(d) - 1; {
    if (d[i] > d[i + 1]) {
      c := d[i + 1]
      d[i + 1] = d[i]
      d[i] = c
      i--
    } else {
      i++
    }
  }
  return d
}

func main() {
  var a, b int
  fmt.Scan(&a)
  var d []int
  for i := 0; i < a; i++ {
    fmt.Scan(&b)
    d = append(d, b)
  }
}
