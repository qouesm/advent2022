package main

import (
	"strings"

	"github.com/qouesm/advent2022/read"
)

// 1 rock
// 2 paper
// 3 scissors
// 0 loss
// 3 draw
// 6 win

func q2a() int {
	var score int
	for _, line := range read.ReadLines("input/2.txt") {
		// m: m[0] their move; m[1] my move
		m := strings.Split(line, " ")
		// blank lines
		if len(m) != 2 {
			continue
		}
		// move played
		switch m[1] {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}
		// outcome
		// loss (not calculated because it adds 0)
		// tie
		if m[0] == "A" && m[1] == "X" ||
			m[0] == "B" && m[1] == "Y" ||
			m[0] == "C" && m[1] == "Z" {
			score += 3
			// win
		} else if m[0] == "A" && m[1] == "Y" ||
			m[0] == "B" && m[1] == "Z" ||
			m[0] == "C" && m[1] == "X" {
			score += 6
		}
	}
	return score
}

func q2b() int {
	var score int
	for _, line := range read.ReadLines("input/2.txt") {
		// m: m[0] their move; m[1] my move
		m := strings.Split(line, " ")
		// blank lines
		if len(m) != 2 {
			continue
		}
		switch m[1] {
		// loss
		case "X":
			switch m[0] {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
			// tie
		case "Y":
			score += 3
			switch m[0] {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
			// win
		case "Z":
			score += 6
			switch m[0] {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}
	}
	return score
}
