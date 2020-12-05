package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
)

type passport struct {
	Byr string `json:"byr"`
	Iyr string `json:"iyr"`
	Eyr string `json:"eyr"`
	Hgt string `json:"hgt"`
	Hcl string `json:"hcl"`
	Ecl string `json:"ecl"`
	Pid string `json:"pid"`
	Cid string `json:"cid"`
}

func (p *passport) isValid() bool {
	return p.Byr != "" && p.Iyr != "" && p.Eyr != "" && p.Hgt != "" && p.Hcl != "" && p.Ecl != "" && p.Pid != ""
}

func (p *passport) reset() {
	p.Byr = ""
	p.Iyr = ""
	p.Eyr = ""
	p.Hgt = ""
	p.Hcl = ""
	p.Ecl = ""
	p.Pid = ""
	p.Cid = ""
}

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	validPassports := 0
	var p passport

	for _, line := range lines {
		if line == "" {
			if p.isValid() {
				fmt.Print("  valid: ")
				_ = json.NewEncoder(os.Stdout).Encode(p)
				validPassports++
			} else {
				fmt.Print("invalid: ")
				_ = json.NewEncoder(os.Stdout).Encode(p)
			}

			p.reset()
			continue
		}

		fields := strings.Split(line, " ")

		for _, field := range fields {
			fields2 := strings.Split(field, ":")

			if len(fields2) != 2 {
				log.Fatalf("fields2 has too many values: %v\n", fields2)
			}

			k, v := fields2[0], fields2[1]

			switch k {
			case "byr":
				p.Byr = v
			case "iyr":
				p.Iyr = v
			case "eyr":
				p.Eyr = v
			case "hgt":
				p.Hgt = v
			case "hcl":
				p.Hcl = v
			case "ecl":
				p.Ecl = v
			case "pid":
				p.Pid = v
			case "cid":
				p.Cid = v
			default:
				log.Fatalf("unknown kv pair: %s = %v\n", k, v)
			}
		}
	}

	if p.isValid() {
		validPassports++
	}

	fmt.Printf("We found %d valid passports\n", validPassports)
}
