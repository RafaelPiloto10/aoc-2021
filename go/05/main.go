package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"math"
)

type Vector struct {
	x int
	y int
}

type LineSegement struct {
	a *Vector
	b *Vector
}

func (s *LineSegement) Intersects(other *LineSegement) bool {
	i1 := Vector{int(math.Min(float64(s.a.x), float64(s.b.x))),
		int(math.Max(float64(s.a.x), float64(s.b.x)))}

	i2 := Vector{int(math.Min(float64(other.a.xi), float64(other.b.x))),
		int(math.Max(float64(other.a.x), float64(other.b.x)))}
	
	ia := Vector{int(math.Max(
		math.Min(float64(s.a.x), float64(s.b.x)),
		math.Min(float64(other.a.x), float64(other.b.x)))),
		int(math.Min(
			math.Max(float64(s.a.x), float64(s.b.x)),
			math.Max(float64(other.a.x), float64(other.b.x))
		))
	}

}

func main() {
	body, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(body), "\n")

	lineSegments := []*LineSegement{}

	for _, line := range data {
		if line == "" {
			continue
		}

		nums := strings.Split(line, " -> ")
		a := numToVector(nums[0])
		b := numToVector(nums[1])
		if a.x != b.x && a.y != b.y {
			continue
		}

		segment := LineSegement{a, b}
		lineSegments = append(lineSegments, &segment)
	}

	count := 0
	for i, s1 := range lineSegments {
		for j, s2 := range lineSegments {
			if i == j {
				continue
			}

			if s1.Intersects(s2) {
				count++
			}
		}
	}

	fmt.Printf("Total intersecting lines: %v", count)
}


func numToVector(input string) *Vector {
	ns := strings.Split(input, ",")
	a, err := strconv.Atoi(ns[0])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(ns[1])
	if err != nil {
		panic(err)
	}

	v := Vector{a, b}
	return &v
}
