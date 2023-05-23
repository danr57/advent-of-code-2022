// Package day6 implements the solution to Day 6 of the Advent of Code 2022.
package day6

type (
	// Day contains the required data to compute the solution.
	Day struct {
		DataStream []string
		InputFile  string
	}
)

// New returns a new instance of Day.
func New() *Day {
	return &Day{}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 6
}
