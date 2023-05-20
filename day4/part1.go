package day4

import "strconv"

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	busyTeams := []*Team{}

	for _, team := range d.Teams {
		if d.checkForTotalOverlap(team) {
			busyTeams = append(busyTeams, team)
		}
	}

	return strconv.Itoa(len(busyTeams)), nil
}
