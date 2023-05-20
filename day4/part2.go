package day4

import "strconv"

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	busyTeams := []*Team{}

	for _, team := range d.Teams {
		if d.checkForAnyOverlap(team) {
			busyTeams = append(busyTeams, team)
		}
	}

	return strconv.Itoa(len(busyTeams)), nil
}
