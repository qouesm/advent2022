package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/qouesm/advent2022/read"
)

func q1a() int {
	var elves []int // contains the total calories calories carried by the ith elf
	for i, e := range strings.Split(read.ReadFile("input/1.txt"), "\n\n") {
		elves = append(elves, 0)
		for _, v := range strings.Split(e, "\n") {
			int, err := strconv.Atoi(v)
			if err != nil {
			}
			elves[i] += int
		}
	}
	sort.Ints(elves)
	return elves[len(elves)-1]
}

func q1b() int {
	var elves []int // contains the total calories calories carried by the ith elf
	for i, e := range strings.Split(read.ReadFile("input/1.txt"), "\n\n") {
		elves = append(elves, 0)
		for _, v := range strings.Split(e, "\n") {
			int, err := strconv.Atoi(v)
			if err != nil {
			}
			elves[i] += int
		}
	}
	sort.Ints(elves)
	var sum int
	for i := 1; i <= 3; i++ {
		sum += elves[len(elves)-i]
	}
  return sum
}
