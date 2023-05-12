package d1

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type (
	Elf struct {
		Foods    []int
		Calories int
	}
)

var Elves []Elf

func Part1() int {
	ans := 0
	for _, n := range Elves {
		if n.Calories > ans {
			ans = n.Calories
		}
	}
	return ans
}

func ReadAndSplit(r string) int {
	input, err := os.Open(filepath.Clean(r))
	if err != nil {
		log.Fatalf("Error reading file: %v \n Error: %v", r, err)
	}

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	elf := Elf{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			Elves = append(Elves, elf)
			elf = Elf{}

			continue
		}

		food, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error conversting line to int: %v", err)
		}

		elf.Calories += food

		elf.Foods = append(elf.Foods, food)

	}

	return len(Elves)
}
