package main

import (
	"flag"
	"fmt"
	"os"

	"encoding/xml"

	"gopkg.in/yaml.v3"
)

type Valc struct {
	Vals string `xml:"Valute"`
}

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
	f2, _ := os.ReadFile(conf.Inp)
	var vc Valc
	_ = xml.Unmarshal(f2, vc)
	fmt.Println(vc)
}
