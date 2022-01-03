package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatalf("unable to read file %v", err)
	}

	data := string(body)
	lines := strings.Split(data, "\n")

	o2 := make([]string, len(lines))
	co2 := make([]string, len(lines))

	copy(o2, lines)
	copy(co2, lines)

	for i := 0; i < len(o2[0]); i++ {
		freq := getColFreq(o2, i, true)
		o2 = removeWithBitAt(o2, i, freq)
	}
	for i := 0; i < len(co2[0]); i++ {
		freq := getColFreq(co2, i, false)
		co2 = removeWithBitAt(co2, i, freq)
	}

	fmt.Println(len(o2))
	fmt.Println(co2)

	o2Byte, err := strconv.ParseUint(o2[0], 2, 64)
	if err != nil {
		log.Fatalf("unable to parse %v to bit; %v", o2[0], err)
	}

	co2Byte, err := strconv.ParseUint(co2[0], 2, 64)
	if err != nil {
		log.Fatalf("unable to parse %v to bit; %v", co2[0], err)
	}

	rating := o2Byte * co2Byte

	fmt.Printf("%v\n", rating)
}

func removeWithBitAt(s []string, index int, bit byte) []string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == "" || s[i] == " " {
			s = remove(s, i)
			continue
		}

		if s[i][index] != bit && len(s) != 1 {
			s = remove(s, i)
		}
	}

	return s
}

func getColFreq(s []string, index int, getMax bool) byte {
	colCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}

		if s[i][index] == '0' {
			colCount -= 1
		} else {
			colCount += 1
		}
	}

	if colCount < 0 {
		if getMax {
			return '0'
		}

		return '1'
	} else {
		if getMax {
			return '1'
		}
		return '0'
	}
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func part1() {
	body, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatalf("unable to read file %v", err)
	}

	data := string(body)
	lines := strings.Split(data, "\n")

	gammaRate := ""
	epsilonRate := ""

	for j := 0; j < len(lines[0]); j++ {
		colCount := 0
		for i := 0; i < len(lines); i++ {
			if lines[i] == "" {
				continue
			}

			fmt.Println(lines[i])
			if lines[i][j] == '0' {
				colCount -= 1
			} else {
				colCount += 1
			}
		}

		if colCount < 0 {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	gammaByte, err := strconv.ParseUint(gammaRate, 2, 64)
	if err != nil {
		log.Fatalf("unable to parse %v to bit; %v", gammaRate, err)
	}

	epsilonByte, err := strconv.ParseUint(epsilonRate, 2, 64)
	if err != nil {
		log.Fatalf("unable to parse %v to bit; %v", epsilonRate, err)
	}

	power := gammaByte * epsilonByte

	fmt.Printf("%v\n", power)
}
