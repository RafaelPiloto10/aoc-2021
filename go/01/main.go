package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func main() {
	body, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := strings.Split(string(body), "\n")
	prev, err := strconv.Atoi(data[0])

	if err != nil {
		log.Fatalf("could not parse %v to int", data[0])
	}

	count := 0

	for i := 0; i < len(data) - 3; i++{
		current, err := strconv.Atoi(data[i + 3])

		if err != nil {
			continue
		}

		if (current > prev) {
			count++
		}
		prev, err = strconv.Atoi(data[i + 1])
	}

	fmt.Println(count)
}
