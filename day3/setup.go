package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	readFile, err := os.Open(d.InputFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer func() {
		if err := readFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		bag := &Bag{}

		line := fileScanner.Text()
		splitPoint := len(line) / 2 //nolint:gomnd // halving a value, leave me alone

		items := strings.Split(line, "")
		bag.AllItems = mapItems(items)
		bag.Compartments[0].MappedItems = mapItems(items[:splitPoint])
		bag.Compartments[1].MappedItems = mapItems(items[splitPoint:])

		d.Bags = append(d.Bags, bag)
	}

	d.makeGroups(d.Bags)

	return nil
}

func mapItems(items []string) map[string]int {
	mappedItems := make(map[string]int)

	for _, item := range items {
		mappedItems[item]++
	}

	return mappedItems
}
