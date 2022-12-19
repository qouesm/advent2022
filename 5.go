package main

import (
	"strconv"
	"strings"

	"github.com/qouesm/advent2022/read"
)

func q5a() string {
	file := read.ReadFile("input/5.txt")
	splitFile := strings.Split(file, "\n\n")
	topLines := strings.Split(splitFile[0], "\n")
	bottomLines := strings.Split(splitFile[1], "\n")

	// parse stacks

	var nStacks int = len(strings.Fields(topLines[len(topLines)-1]))
	stacks := make([]Stack, nStacks)

	var elemXPositions []int
	for x := 1; x < len(topLines[0]); x += 4 {
		elemXPositions = append(elemXPositions, x)
	}

	// reverse iterate through stacks, excluding numbered row
	for i := len(topLines) - 2; i >= 0; i-- {
		line := topLines[i]
		// move between each (potential) element
		for i, x := range elemXPositions {
			if line[x] == ' ' {
				continue
			}
			stacks[i].Push(rune(line[x]))
		}
	}

	// rearrangement

	for _, inst := range bottomLines {
		if inst == "" {
			continue
		}
		fields := strings.Fields(inst)
		n, _ := strconv.Atoi(fields[1])
		a, _ := strconv.Atoi(fields[3])
		b, _ := strconv.Atoi(fields[5])
		for i := 0; i < n; i++ {
			stacks[b-1].Push(stacks[a-1].Pop())
		}
	}

	tops := ""
	for _, stack := range stacks {
		tops += string(stack.Pop())
	}

	return tops
}

func q5b() string {
	file := read.ReadFile("input/5.txt")
	splitFile := strings.Split(file, "\n\n")
	topLines := strings.Split(splitFile[0], "\n")
	bottomLines := strings.Split(splitFile[1], "\n")

	// parse stacks

	var nStacks int = len(strings.Fields(topLines[len(topLines)-1]))
	stacks := make([]Stack, nStacks)

	var elemXPositions []int
	for x := 1; x < len(topLines[0]); x += 4 {
		elemXPositions = append(elemXPositions, x)
	}

	// reverse iterate through stacks, excluding numbered row
	for i := len(topLines) - 2; i >= 0; i-- {
		line := topLines[i]
		// move between each (potential) element
		for i, x := range elemXPositions {
			if line[x] == ' ' {
				continue
			}
			stacks[i].Push(rune(line[x]))
		}
	}

	// rearrangement

	for _, inst := range bottomLines {
		if inst == "" {
			continue
		}
		fields := strings.Fields(inst)
		n, _ := strconv.Atoi(fields[1])
		a, _ := strconv.Atoi(fields[3])
		b, _ := strconv.Atoi(fields[5])
		hold := new(Stack)
		for i := 0; i < n; i++ {
			hold.Push(stacks[a-1].Pop())
		}
		for i := 0; i < n; i++ {
			stacks[b-1].Push(hold.Pop())
		}
	}

	tops := ""
	for _, stack := range stacks {
		tops += string(stack.Pop())
	}

	return tops
}

type item struct {
	value rune
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value rune) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value rune) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}

	return ' '
}
