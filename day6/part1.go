package day6

import (
	"strconv"
)

const (
	markerLength = 4
)

func (d *Day) getDataSegment(index int, length int) []string {
	var segment []string

	if len(d.DataStream[index:]) < length {
		segment = d.DataStream[index:]
	} else {
		segment = d.DataStream[index : index+length]
	}

	return segment
}

// unique returns true if all the characters in the segment are unique.
func unique(segment []string) bool {
	seen := make(map[string]bool)

	for _, c := range segment {
		if seen[c] {
			return false
		}

		seen[c] = true
	}

	return true
}

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	for index := range d.DataStream {
		segment := d.getDataSegment(index, markerLength)
		if unique(segment) {
			return strconv.Itoa(index + markerLength), nil
		}
	}

	return "", nil
}
