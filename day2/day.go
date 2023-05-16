// Package day2 implements the solution to Day 2 of the Advent of Code 2022.
package day2

type (
	// Day is the implementation of Day 2.
	Day struct {
		InputFile  string
		Part1Score int
		Part2Score int
		Rounds     []*Round
	}
)

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		InputFile: "./day2/input.txt",
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 2
}
