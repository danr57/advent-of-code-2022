package day3

import (
	"strconv"
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	badgePriority := 0

	for _, group := range d.Groups {
		badgePriority += d.findBadge(group)
	}

	return strconv.Itoa(badgePriority), nil
}
