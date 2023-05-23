package day6

import (
	"strconv"
)

const (
	messageLength = 14
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	for index := range d.DataStream {
		segment := d.getDataSegment(index, messageLength)
		if unique(segment) {
			return strconv.Itoa(index + messageLength), nil
		}
	}

	return "", nil
}
