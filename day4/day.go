// Package day4 implements the solution to Day 4 of the Advent of Code 2022.
package day4

type (
	Cleaner struct {
		assignments []int
		min         int
		max         int
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

// Checks that all of a Cleaner's assignments overlap with their partner's assignments
func (d *Day) checkForTotalOverlap(team *Team) bool {
	cleaner1min := team.Cleaners[0].assignments[0]
	cleaner1max := team.Cleaners[0].assignments[len(team.Cleaners[0].assignments)-1]
	cleaner2min := team.Cleaners[1].assignments[0]
	cleaner2max := team.Cleaners[1].assignments[len(team.Cleaners[1].assignments)-1]

	return (cleaner2min >= cleaner1min && cleaner2max <= cleaner1max) ||
		(cleaner1min >= cleaner2min && cleaner1max <= cleaner2max)
}

// Checks whether ANY of 2 Cleaners' assignments overlap
func (d *Day) checkForAnyOverlap(team *Team) bool {
	cleaner1min := team.Cleaners[0].assignments[0]
	cleaner1max := team.Cleaners[0].assignments[len(team.Cleaners[0].assignments)-1]
	cleaner2min := team.Cleaners[1].assignments[0]
	cleaner2max := team.Cleaners[1].assignments[len(team.Cleaners[1].assignments)-1]

	return (cleaner1min >= cleaner2min && cleaner1min <= cleaner2max) ||
		(cleaner2min >= cleaner1min && cleaner2min <= cleaner1max)

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
