package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	yml := flag.String("config", "", "way to config")
	flag.Parse()
	fmt.Println(*yml)
	f, _ := os.Open(*yml)
	buf := make([]byte, 1000)
	for {
		d, err := f.Read(buf)
		if (err != nil) {
			break
		}
		fmt.Println(string(buf[:d]))
	}
}
