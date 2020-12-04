package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

var (
	hairColorRegex *regexp.Regexp
	pidRegex       *regexp.Regexp
)

func init() {
	re, err := regexp.Compile("^#[0-9a-f]{6}$")
	if err != nil {
		log.Fatalf("error compiling regex: %v\n", err)
	}

	hairColorRegex = re

	re2, err := regexp.Compile("^[0-9]{9}$")
	if err != nil {
		log.Fatalf("error compiling regex: %v\n", err)
	}

	pidRegex = re2
}

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
	if p.Byr == "" || p.Iyr == "" || p.Eyr == "" || p.Hgt == "" || p.Hcl == "" || p.Ecl == "" || p.Pid == "" {
		return false
	}

	byr, err := strconv.Atoi(p.Byr)
	if err != nil {
		return false
	}

	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr, err := strconv.Atoi(p.Iyr)
	if err != nil {
		return false
	}

	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, err := strconv.Atoi(p.Eyr)
	if err != nil {
		return false
	}

	if eyr < 2020 || eyr > 2030 {
		return false
	}

	if strings.Contains(p.Hgt, "cm") {
		hgt := strings.TrimSuffix(p.Hgt, "cm")
		h, err := strconv.Atoi(hgt)
		if err != nil {
			return false
		}

		if h < 150 || h > 193 {
			return false
		}
	} else if strings.Contains(p.Hgt, "in") {
		hgt := strings.TrimSuffix(p.Hgt, "in")
		h, err := strconv.Atoi(hgt)
		if err != nil {
			return false
		}

		if h < 59 || h > 76 {
			return false
		}
	} else {
		return false
	}

	switch p.Ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		return false
	}

	return hairColorRegex.MatchString(p.Hcl) && pidRegex.MatchString(p.Pid)
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
	lines := getTheStuff()
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

func getTheStuff() []string {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var results []string

	for ; scanner.Scan(); {
		results = append(results, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return results
}
