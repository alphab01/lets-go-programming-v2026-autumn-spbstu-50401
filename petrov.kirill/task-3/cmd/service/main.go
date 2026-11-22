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
	d, _ := os.ReadFile(*yml)
	fmt.Println(string(d))
}
