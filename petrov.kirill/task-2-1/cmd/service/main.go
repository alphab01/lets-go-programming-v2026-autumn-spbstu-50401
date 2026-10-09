package main

import (
	"fmt"
)

func main() {
	var num, kol int
	_, err := fmt.Scan(&num, &kol)
	if err != nil {
		fmt.Println(err)
	}

}

/*
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

func gogo(kolichestvo int, bolshemenshe string, add int) {
	var left, right int
	for iter := range kolichestvo {
		_, err := fmt.Scan(&bolshemenshe)
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

func run(num int, kolichestvo int, bolshemenshe string, add int, left int, right int) {
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}
	for range num {
		_, err = fmt.Scan(&kolichestvo)
		if err != nil {
			fmt.Println(err)
		}
		gogo(kolichestvo, bolshemenshe, add)
	}
}

func main() {
	var num int
	var kolichestvo int
	var bolshemenshe string
	var add int
	var left int
	var right int
	run(num, kolichestvo, bolshemenshe, add, left, right)
}
*/
