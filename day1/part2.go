package day1

import (
	"sort"
	"strconv"
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	d.SortElves()

	return strconv.Itoa(
		d.Elves[0].Calories +

			d.Elves[1].Calories +
			d.Elves[2].Calories), nil
}

// SortElves sorts elves by highest calories.
func (d *Day) SortElves() {
	sort.Slice(d.Elves, func(i, j int) bool {
		return d.Elves[i].Calories > d.Elves[j].Calories
	})
}
