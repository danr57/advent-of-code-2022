// Package d1 implements the solution to Day 1 of the Advent of Code 2022.
package d1

type (
	// Elf contains the information about an elf.
	// Foods contain the calorie count for each of the food items they have.
	// Calories is the total calorie count for all the food items they have.
	Elf struct {
		Foods    []int
		Calories int
	}

	// Day contains our data for the day and the file to get it from.
	Day struct {
		Elves     []*Elf
		InputFile string
	}
)

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		InputFile: "./d1/input",
	}
}

// Number returns the day number.
//
// nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 1
}
