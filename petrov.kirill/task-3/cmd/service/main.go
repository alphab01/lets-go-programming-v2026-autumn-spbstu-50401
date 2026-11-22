package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Valu struct {
	Numc  int    `xml:"NumCode"`
	Charc string `xml:"CharCode"`
	Nom   int    `xml:"Nominal"`
	Name  string `xml:"Name"`
	Val   string `xml:"Value"`
	Vunt  string `xml:"VunitRate"`
}

type Valc struct {
	Vals []Valu `xml:"Valute"`
}

type Conf struct {
	Inp string `yaml:"input-file"`
	Out string `yaml:"output-file"`
}

type jsonsave struct {
	Numc  int     `json:"num_code"`
	Charc string  `json:"char_code"`
	Val   float64 `json:"value"`
}

func zapataya(s string) string {
	s1 := ""
	for i := range len(s) {
		if string(s[i]) != "," {
			s1 += string(s[i])
		} else {
			s1 += "."
		}
	}

	return s1
}

func sortvalu(d []Valu) []Valu {
	for i := 0; i < len(d); i++ {
		for j := range len(d) - 1 {
			a, err := strconv.ParseFloat(zapataya(d[j].Val), 64)
			if err != nil {
				panic("Parse a !work")
			}
			b, err := strconv.ParseFloat(zapataya(d[j+1].Val), 64)
			if err != nil {
				panic("Parse b !work")
			}
			if a < b {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
	}

	return d
}

func vtj1(aaa Valu) jsonsave {
	ddd, err := strconv.ParseFloat(zapataya(aaa.Val), 64)
	if err != nil {
		panic("parsing to float not work")
	}
	q := jsonsave{Numc: aaa.Numc, Charc: aaa.Charc, Val: ddd}

	return q
}

func VtJ(vals []Valu) []jsonsave {
	var js []jsonsave
	for i := range len(vals) {
		js = append(js, vtj1(vals[i]))
	}
	return js
}

func main() {
	yml := flag.String("config", "", "way to config")
	flag.Parse()
	f, err := os.ReadFile(*yml)
	if err != nil {
		panic("Read config not works")
	}
	var conf Conf
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		panic(err)
	}
	f2, err := os.ReadFile(conf.Inp)
	if err != nil {
		panic(err)
	}
	var vc Valc
	dec := xml.NewDecoder(bytes.NewReader(f2))
	dec.CharsetReader = charset.NewReaderLabel
	err = dec.Decode(&vc)
	if err != nil {
		panic("Decoder of xml lists works bad")
	}
	vc.Vals = sortvalu(vc.Vals)
	js := VtJ(vc.Vals)
	aaa, err := json.Marshal(js)
	if err != nil {
		panic("Perevod v jeysons not work")
	}
	err = os.MkdirAll(filepath.Dir(conf.Out), 0o777)
	if err != nil {
		panic("mkdir err")
	}
	f3, err := os.Create(conf.Out)
	if err != nil {
		panic("create err")
	}
	_, err = f3.Write(aaa)
	if err != nil {
		panic("write err")
	}
	err = f3.Close()
	if err != nil {
		panic("close f3 err")
	}
}
