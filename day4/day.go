// Package day4 implements the solution to Day 4 of the Advent of Code 2022.
package day4

type (
	// Day is the implementation of Day 4.
	Day struct{
		InputFile string
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
	return 4
}

