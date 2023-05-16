// Package day5 implements the solution to Day 5 of the Advent of Code 2022.
package day5

type (
	// Move crates from one stack to another.
	Move struct {
		crates int
		from   int
		to     int
	}

	// Day contains the stacks and moves for a given day.
	Day struct {
		Stacks    [][]string
		Moves     []Move
		InputFile string
	}
)

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		InputFile: "./day5/input.txt",
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 5
}
