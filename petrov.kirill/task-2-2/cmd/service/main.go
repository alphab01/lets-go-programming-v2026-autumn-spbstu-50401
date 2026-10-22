package main

import (
	"fmt"
)

type bludo []int

func (b bludo) Len() int {

	return len(b)
}

func (b bludo) get(i int) int {

	return b[i]
}

func (b *bludo) add(a int) {
	*b = append(*b, a)
}

func (b *bludo) fix(i int, a int) {
	(*b)[i] = a
}

func bsort(d bludo) bludo {
	for range d.Len() {
		for j := range d.Len()-1 {
			if d.get(j) < d.get(j+1) {
				c := d.get(j + 1)
				d.fix(j+1, d.get(j))
				d.fix(j, c)
			}
		}
	}

	return d
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println(err)

		return
	}
	d := bludo{}
	for range a {
		fmt.Scan(&b)
		d.add(b)
	}
	fmt.Scan(&b)
	d = bsort(d)
	fmt.Println(d.get(b - 1))
}
