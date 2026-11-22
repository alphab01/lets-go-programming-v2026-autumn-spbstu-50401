package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Valu struct {
	Numc string `xml:"NumCode"`
	Charc string `xml:"CharCode"`
	Nom string `xml:"Nominal"`
	Name string `xml:"Name"`
	Val string `xml:"Value"`
	Vunt string `xml:"VunitRate"`
}

type Valc struct {
	Vals []Valu `xml:"Valute"`
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
	dec := xml.NewDecoder(bytes.NewReader(f2))
	dec.CharsetReader = charset.NewReaderLabel
	_ = dec.Decode(&vc)
	fmt.Println(vc.Vals[0])
}
