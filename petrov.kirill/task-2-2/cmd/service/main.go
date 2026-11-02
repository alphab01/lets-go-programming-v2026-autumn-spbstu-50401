package main

import (
	"fmt"
  "container/heap"
)

type bludo []int



func (bld bludo) Len() int {
	return len(bld)
}

func (bld bludo) Less(j int, i int) bool {
	return bld[i] < bld[j]
}

func (bld bludo) Swap(i int, j int) {
	swp := bld[i]
  bld[i] = bld[j]
  bld[j] = swp
}

func main() {
	var aaa, bbb int

	_, err := fmt.Scan(&aaa)
	if err != nil {
		fmt.Println(err)

		return
	}

	dat := bludo{}

	for range aaa {
		_, err = fmt.Scan(&bbb)
		if err != nil {
			fmt.Println(err)

			return
		}

		dat.add(bbb)
	}

	_, err = fmt.Scan(&bbb)
	if err != nil {
		fmt.Println(err)

		return
	}

	dat = bsort(dat)
	fmt.Println(dat.get(bbb - 1))
}
