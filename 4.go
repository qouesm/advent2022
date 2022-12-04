package main

import (
	"strconv"
	"strings"

	"github.com/qouesm/advent2022/read"
)

func q4a() int {
	var count int
	for _, line := range read.ReadLines("input/4.txt") {
		var id []int // contains the four numbers on each line
		for _, v := range strings.Split(line, ",") {
			temp := strings.Split(v, "-")
			for _, w := range temp {
				int, err := strconv.Atoi(w)
				if err != nil {
				}
				id = append(id, int)
			}
		}
		if id[0] >= id[2] && id[1] <= id[3] ||
			id[2] >= id[0] && id[3] <= id[1] {
			count++
		}
	}
	return count
}

func q4b() int {
	var count int
	for _, line := range read.ReadLines("input/4.txt") {
		var id []int // contains the four numbers on each line
		for _, v := range strings.Split(line, ",") {
			temp := strings.Split(v, "-")
			for _, w := range temp {
				int, err := strconv.Atoi(w)
				if err != nil {
				}
				id = append(id, int)
			}
		}
    stop := false
    for i := id[0]; i <= id[1]; i++ {
      for j := id[2]; j <= id[3]; j++ {
        if i == j {
          count++
          stop = true
          break
        }
      }
      if stop {
        break
      }
    }
	}
	return count
}

