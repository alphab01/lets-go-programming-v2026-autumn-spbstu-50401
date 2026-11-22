package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Valu struct {
	Numc int `xml:"NumCode"`
	Charc string `xml:"CharCode"`
	Nom int `xml:"Nominal"`
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

type jsonsave struct {
	Numc int
	Charc string
	Val float64	
}

func zapataya(s string) string {
	s1 := ""
	for i := 0; i < len(s); i++ {
		if (string(s[i]) != ",") {
			s1 += string(s[i])
		} else {
			s1 += "."
		}
	}
	return s1
}

func sortvalu(d []Valu) []Valu {
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d) - 1; j++ {
			a, _ := strconv.ParseFloat(zapataya(d[j].Val), 64)
			b, _ := strconv.ParseFloat(zapataya(d[j + 1].Val), 64)
			if (a < b) {
				d[j], d[j + 1] = d[j + 1], d[j]
			}
		}
	}
	return d
}

func vtj1(aaa Valu) jsonsave {
	ddd, _ := strconv.ParseFloat(aaa.Val, 64)
	q := jsonsave{Numc: aaa.Numc, Charc: aaa.Charc, Val: ddd}
	return q
}

func VtJ(vals []Valu) []jsonsave {
	var js []jsonsave
	for i := 0; i < len(vals); i++ {
		js = append(js, vtj1(vals[i]))
	}
	return js
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
	vc.Vals = sortvalu(vc.Vals)
	fmt.Println(vc.Vals[0])
	fmt.Println(vc.Vals[len(vc.Vals) - 1])
	js := VtJ(vc.Vals)
	fmt.Println(js[0].Charc)
}
