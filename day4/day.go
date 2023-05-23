// Package day4 implements the solution to Day 4 of the Advent of Code 2022.
package day4

type (
	Cleaner struct {
		min int
		max int
	}

	Team struct {
		Cleaners []*Cleaner
	}

	// Day is the implementation of Day 4.
	Day struct {
		Teams     []*Team
		InputFile string
	}
)

// Checks that all of a Cleaner's assignments overlap with their partner's assignments.
func (d *Day) checkForTotalOverlap(team *Team) bool {
	cleaner1 := team.Cleaners[0]
	cleaner2 := team.Cleaners[1]

	return (cleaner2.min >= cleaner1.min && cleaner2.max <= cleaner1.max) ||
		(cleaner1.min >= cleaner2.min && cleaner1.max <= cleaner2.max)
}

// Checks whether ANY of 2 Cleaners' assignments overlap.
func (d *Day) checkForAnyOverlap(team *Team) bool {
	cleaner1 := team.Cleaners[0]
	cleaner2 := team.Cleaners[1]

	return (cleaner1.min >= cleaner2.min && cleaner1.min <= cleaner2.max) ||
		(cleaner2.min >= cleaner1.min && cleaner2.min <= cleaner1.max)
}

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		InputFile: "./day4/input.txt",
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 4
}
