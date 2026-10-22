package main

import (
  "fmt"
)

type Interface interface {
  sort.Interface
  Push(x interface{})
  Pop() interface{}
}

func bsort(d []int) []int {
  for i := 0; i < len(d); i++ {
    for j := 0; j < len(d) - 1; j++ {
      if (d[j] < d[j + 1]) {
        c := d[j + 1]
        d[j + 1] = d[j]
        d[j] = c
      }
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
  fmt.Scan(&b)
  d = bsort(d)
  fmt.Println(d[b - 1])
}
