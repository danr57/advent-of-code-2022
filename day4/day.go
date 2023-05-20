// Package day4 implements the solution to Day 4 of the Advent of Code 2022.
package day4

type (
	Cleaner struct {
		assignments []int
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

func (d *Day) checkForTotalOverlap(team *Team) bool {
	cleaner1min := team.Cleaners[0].assignments[0]
	cleaner1max := team.Cleaners[0].assignments[len(team.Cleaners[0].assignments)-1]
	cleaner2min := team.Cleaners[1].assignments[0]
	cleaner2max := team.Cleaners[1].assignments[len(team.Cleaners[1].assignments)-1]

	if cleaner1min >= cleaner2min && cleaner1max <= cleaner2max {
		return true

	}
	if cleaner2min >= cleaner1min && cleaner2max <= cleaner1max {
		return true
	}

	return false
}

func (d *Day) checkForAnyOverlap(team *Team) bool {
	cleaner1 := team.Cleaners[0]
	cleaner2 := team.Cleaners[1]

	if cleaner1.assignments[0] >= cleaner2.assignments[0] {
		return true
	}

	return false
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
