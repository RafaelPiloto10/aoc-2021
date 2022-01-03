package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

type Vector struct {
	horizontal int
	depth int
	aim int
}

func (v *Vector) add(other *Vector) {
	v.horizontal += other.horizontal
	v.depth += other.depth
	v.aim += other.aim
}

func main() {
	body, err := ioutil.ReadFile("./input.txt")
	
	if err != nil {
		log.Fatalf("unable to read file %v", err)	
	}

	data := string(body)
	lines := strings.Split(data, "\n") 

	pos := Vector{0, 0, 0}

	for _, line := range(lines) {
		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")
		command := tokens[0]
		dist, _ := strconv.Atoi(tokens[1])
		res := CommandToPos(command, dist, &pos)
		pos.add(&res)
	}
	
	fmt.Println(pos.depth * pos.horizontal)
}

func CommandToPos(command string, dist int, pos *Vector) Vector {
	if (command == "forward") {
		return Vector{dist, dist * pos.aim, 0}
	} else if(command == "down") {
		return Vector{0, 0, dist}
	} else if(command == "up") {
		return Vector{0, 0, -dist}
	}

	return Vector{0, 0, 0}
}
