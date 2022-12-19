package main

import (
	"github.com/qouesm/advent2022/read"
)

func q6a() int {
	input := read.ReadFile("input/6.txt")
	// initial sliding carriage
	carriage := input[:4]

	// find valid
	for i, v := range input[4:] {
		carriage = carriage[1:] + string(v)
		match := false
		for i := 0; i < len(carriage); i++ {
			for j := i + 1; j < len(carriage); j++ {
				if carriage[i] == carriage[j] {
					match = true
					break
				}
			}
		}
		if !match {
			return i + 5 // +1 for one-indexing; +4 to get to end of carriage
		}
	}

	return -1
}

func q6b() int {
	input := read.ReadFile("input/6.txt")
	// initial sliding carriage
	carriage := input[:14]

	// find valid
	for i, v := range input[14:] {
		carriage = carriage[1:] + string(v)
		match := false
		for i := 0; i < len(carriage); i++ {
			for j := i + 1; j < len(carriage); j++ {
				if carriage[i] == carriage[j] {
					match = true
					break
				}
			}
		}
		if !match {
			return i + 15 // +1 for one-indexing; +15 to get to end of carriage
		}
	}

	return -1
}
