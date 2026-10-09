package main

import (
	"fmt"
)

func ismore(str string) bool {
	return string(str[0]) == ">"
}

const (
	mini = 15
	maxi = 30
)

func printt(left int, right int) {
	if right > maxi {
		right = maxi
	}
	if left < mini {
		left = mini
	}
	if left <= right {
		fmt.Println(left)
	} else {
		fmt.Println(-1)
	}
}

func run(Num int, Kolichestvo int, bolshemenshe string, add int, left int, right int) {
  _, err := fmt.Scan(&Num)
  if err != nil {
    fmt.Println(err)
  }
  for _ = range Num {
    _, err = fmt.Scan(&Kolichestvo)
    if err != nil {
      fmt.Println(err)
    }
    for iter := range Kolichestvo {
      _, err = fmt.Scan(&bolshemenshe)
      if err != nil {
        fmt.Println(err)
      }
      _, err = fmt.Scan(&add)
      if err != nil {
        fmt.Println(err)
      }
      if iter == 0 {
        if ismore(bolshemenshe) {
          left = add
          right = maxi
        } else {
          right = add
          left = mini
        }
      } else {
        if ismore(bolshemenshe) {
          if add > left {
            left = add
          }
        } else {
          if add < right {
            right = add
          }
        }
      }
      printt(left, right)
    }
  }
}

func main() {
	var Num int
	var Kolichestvo int
	var bolshemenshe string
	var add int
	var left int
	var right int
  run(Num, Kolichestvo, bolshemenshe, add, left, right)
}
