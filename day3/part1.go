package day3

import (
	"strconv"
)

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	for _, bag := range d.Bags {
		for item := range bag.Compartments[0].MappedItems {
			if _, exists := bag.Compartments[1].MappedItems[item]; exists {
				d.PrioritySum += d.PriorityMap[item]
			}
		}
	}

	return strconv.Itoa(d.PrioritySum), nil
}
