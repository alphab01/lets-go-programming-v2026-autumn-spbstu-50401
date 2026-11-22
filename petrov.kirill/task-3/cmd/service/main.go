package main

import (
	"flag"
	"fmt"
	"os"
  "gopkg.in/yaml.v3"
)

type Conf struct {
  Inp string `yaml:"input-file"`
  Out string `yaml:"output-file"`
}

func main() {
	yml := flag.String("config", "", "way to config")
	flag.Parse()
	fmt.Println(*yml)
	f, _ := os.ReadFile(*yml)
  var conf Conf
  _ = yaml.Unmarshal(f, &conf)
  fmt.Println(conf.Inp)
}
