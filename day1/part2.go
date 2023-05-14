package day1

import "strconv"

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	return strconv.Itoa(
		d.Elves[0].Calories +

			d.Elves[1].Calories +
			d.Elves[2].Calories), nil
}
