package main

import (
	"flag"
	"fmt"
)

func main() {
	yml := flag.String("config", "", "way to config")
	flag.Parse()
	fmt.Println(*yml)
}
