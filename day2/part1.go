package day2

import (
	"errors"
	"strconv"
)

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	if d.Score == 0 {
		err := errors.New(`d.Score is nil`)
		return "", err
	}
	return strconv.Itoa(d.Score), nil
}
