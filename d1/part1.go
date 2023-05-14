package d1

import "strconv"

func (d *Day) Part1() (string, error) {
	ans := 0
	for _, n := range d.Elves {
		if n.Calories > ans {
			ans = n.Calories
		}
	}
	return strconv.Itoa(ans), nil
}
