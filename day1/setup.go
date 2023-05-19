package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	input, err := os.Open(d.InputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v \n Error: %v", input, err)

		return err
	}

	scanner := bufio.NewScanner(input)

	scanner.Split(bufio.ScanLines)

	elf := &Elf{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			d.Elves = append(d.Elves, elf)
			elf = &Elf{}

			continue
		}

		food, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error conversting line to int: %v", err)
		}

		elf.Calories += food

		elf.Foods = append(elf.Foods, food)
	}

	d.Elves = append(d.Elves, elf)

	return nil
}
