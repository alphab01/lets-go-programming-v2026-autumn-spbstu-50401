package main

import (
	"fmt"
)

func ismore(str string) bool {
	if string(str[0]) == ">" {
		return true
	}
	return false
}

func printt(left int, right int) {
	if right > 30 {
		right = 30
	}
	if left < 15 {
		left = 15
	}
	if left <= right {
		fmt.Println(left)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	var Num int
	var Kolichestvo int
	var bolshemenshe string
	var add int
	var left int
	var right int
	_, err := fmt.Scan(&Num)
	if err != nil {
		fmt.Println(err)
	}
	for _ = range Num {
		fmt.Scan(&Kolichestvo)
		for j := range Kolichestvo {
			_, err = fmt.Scan(&bolshemenshe)
			if err != nil {
				fmt.Println(err)
			}
			_, err = fmt.Scan(&add)
			if err != nil {
				fmt.Println(err)
			}
			if j == 0 {
				if ismore(bolshemenshe) {
					left = add
					right = 30
				} else {
					right = add
					left = 15
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
