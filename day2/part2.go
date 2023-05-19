package day2

import (
	"strconv"
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	return strconv.Itoa(d.Part2Score), nil
}
