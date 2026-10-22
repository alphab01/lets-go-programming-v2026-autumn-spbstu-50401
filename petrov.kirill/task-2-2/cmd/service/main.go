package main

import (
	"fmt"
)

type bludo []int

func (bld *bludo) Len() int {
	return len(*bld)
}

func (bld *bludo) get(i int) int {
	return (*bld)[i]
}

func (bld *bludo) add(addd int) {
	*bld = append(*bld, addd)
}

func (bld *bludo) fix(iii int, addd int) {
	(*bld)[iii] = addd
}

func bsort(dat bludo) bludo {
	for range dat.Len() {
		for jjj := range dat.Len() - 1 {
			if dat.get(jjj) < dat.get(jjj+1) {
				ccc := dat.get(jjj + 1)
				dat.fix(jjj+1, dat.get(jjj))
				dat.fix(jjj, ccc)
			}
		}
	}

	return dat
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
