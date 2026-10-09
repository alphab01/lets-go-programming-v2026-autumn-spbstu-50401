package main

import (
  "fmt"
)

func ismore(s string) bool {
  if (string(s[0]) == ">") {
    return true
  }
  return false
}

func printt(l int, r int) {
  if (r > 30) {
    r = 30
  }
  if (l < 15) {
    l = 15
  }
  if (l <= r) {
    fmt.Println(l)
  } else {
    fmt.Println(-1)
  }
}

func main() {
  var N int
  var Kolichestvo int
  var bolshemenshe string
  var add int
  var l int
  var r int
  _, err := fmt.Scan(&N)
  if (err != nil) {
    fmt.Println(err)
  }
  for i := 0; i < N; i++ {
    fmt.Scan(&Kolichestvo)
    for j := 0; j < Kolichestvo; j++ {
      _, err = fmt.Scan(&bolshemenshe)
      if (err != nil) {
        fmt.Println(err)
      }
      _, err = fmt.Scan(&add)
      if (err != nil) {
        fmt.Println(err)
      }
      if (j == 0) {
        if (ismore(bolshemenshe)) {
          l = add
          r = 30
        } else {
          r = add
          l = 15
        }
      } else {
        if (ismore(bolshemenshe)) {
          if (add > l) {
            l = add
          }
        } else {
          if (add < r) {
            r = add
          }
        }
      }
      printt(l, r)
    }
  }
}
