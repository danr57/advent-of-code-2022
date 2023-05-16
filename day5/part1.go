package day5

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	stacks := copyStacks(d.Stacks)

	for _, m := range d.Moves {
		var crates []string
		crates, stacks[m.from] = pop(stacks[m.from], m.crates)
		stacks[m.to] = push(stacks[m.to], reverseSlice(crates))
	}

	var ans string

	for _, s := range stacks {
		ans += s[len(s)-1]
	}

	return ans, nil
}
