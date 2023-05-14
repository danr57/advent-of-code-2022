package day1

import "strconv"

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	ans := 0
	for _, n := range d.Elves {
		if n.Calories > ans {
			ans = n.Calories
		}
	}
	return strconv.Itoa(ans), nil
}
