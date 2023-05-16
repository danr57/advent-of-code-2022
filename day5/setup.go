package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func pop(stack []string, count int) (popped []string, newStack []string) {
	return stack[len(stack)-count:], stack[:len(stack)-count]
}
func push(stack []string, vals []string) (newStack []string) {
	return append(stack, vals...)
}

func copyStack(stack []string) []string {
	newStack := make([]string, len(stack))

	copy(newStack, stack)

	return newStack
}

func copyStacks(stacks [][]string) [][]string {
	newStacks := make([][]string, len(stacks))

	for i, stack := range stacks {
		newStacks[i] = copyStack(stack)
	}

	return newStacks
}

func convertStacks(stackLines []string) [][]string {
	stackIndexes := []int{}

	for i, char := range stackLines[len(stackLines)-1] {
		if string(char) != " " {
			stackIndexes = append(stackIndexes, i)
		}
	}

	stacks := make([][]string, len(stackIndexes))

	stackLines = stackLines[:len(stackLines)-1]

	for _, stackLine := range stackLines {
		for i, stackIndex := range stackIndexes {
			val := string(stackLine[stackIndex])
			if val != " " {
				stacks[i] = append(stacks[i], val)
			}
		}
	}

	for _, stack := range stacks {
		reverseSlice(stack)
	}

	return stacks
}

func convertMoves(moveLines []string) ([]Move, error) {
	moves := []Move{}

	for _, move := range moveLines {
		move = strings.ReplaceAll(move, "move ", "")
		move = strings.ReplaceAll(move, "from ", "")
		move = strings.ReplaceAll(move, "to ", "")
		moveNums := []int{}

		for _, c := range strings.Split(move, " ") {
			n, err := strconv.Atoi(c)
			if err != nil {
				return nil, fmt.Errorf("could not convert to int: %w", err)
			}

			moveNums = append(moveNums, n)
		}

		moves = append(moves, Move{
			crates: moveNums[0],
			from:   moveNums[1] - 1,
			to:     moveNums[2] - 1,
		})
	}

	return moves, nil
}

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	readFile, err := os.Open(d.InputFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer func() {
		if readErr := readFile.Close(); readErr != nil {
			log.Fatal(readErr)
		}
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	stackLines := []string{}
	moveLines := []string{}
	spacerFound := false

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			spacerFound = true

			continue
		}

		if spacerFound {
			moveLines = append(moveLines, line)
		} else {
			stackLines = append(stackLines, line)
		}
	}

	d.Stacks = convertStacks(stackLines)

	d.Moves, err = convertMoves(moveLines)
	if err != nil {
		return err
	}

	return nil
}
