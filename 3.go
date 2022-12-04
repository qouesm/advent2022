package main

import (
	"unicode"

	"github.com/qouesm/advent2022/read"
)

func q3a() int {
	var sum int
	for _, line := range read.ReadLines("input/3.txt") {
		c := []string{line[:len(line)/2], line[len(line)/2:]} // split line in half
		for _, v := range c[1] {
			if in(c[0], v) {
				// ['A', 'Z'] : [65,  90]
				// ['a', 'z'] : [97, 122]
				if unicode.IsUpper(v) {
					sum += int(v - 38)
				} else {
					sum += int(v - 96)
				}
				break
			}
		}
	}
	return sum
}

func q3b() int {
	lines := read.ReadLines("input/3.txt")
	var sum int
	for i := 0; i < len(lines)-1; i += 3 {
		elves := []string{
			lines[i],
			lines[i+1],
			lines[i+2],
		}

		var firstTwoIntersect []rune
		for _, v := range elves[0] {
			if in(elves[1], v) {
				firstTwoIntersect = append(firstTwoIntersect, v)
			}
		}
		for _, v := range firstTwoIntersect {
			if in(elves[2], v) {
				// ['A', 'Z'] : [65,  90]
				// ['a', 'z'] : [97, 122]
				if unicode.IsUpper(v) {
					sum += int(v - 38)
				} else {
					sum += int(v - 96)
				}
				break
			}
		}
	}
	return sum
}

func in(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}
