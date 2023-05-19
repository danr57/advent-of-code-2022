// Package day3 implements the solution to Day 3 of the Advent of Code 2022.
package day3

type (
	// Bag contains the compartments of a bag.
	Bag struct {
		AllItems     map[string]int
		Compartments [2]Compartment
	}

	// Compartment represents the compartment of a bag and the items it contains.
	Compartment struct {
		MappedItems map[string]int
	}

	// Group represents a group of elves and their respective badge
	Group struct {
		Bags  []*Bag
		badge string
	}

	// Day is the implementation of Day 3.
	Day struct {
		Groups      []*Group
		Bags        []*Bag
		InputFile   string
		PrioritySum int
		PriorityMap map[string]int
		GroupSize   int
	}
)

// CreatePriorityMap returns a map of item to priority.
//
//nolint:gomnd // This is a puzzle.
func CreatePriorityMap() map[string]int {
	return map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
}

func (d *Day) makeGroups(bags []*Bag) {
	group := &Group{}

	for _, bag := range bags {
		group.Bags = append(group.Bags, bag)
		if len(group.Bags) == d.GroupSize {
			d.Groups = append(d.Groups, group)
			group = &Group{}
		}
	}
}

func (d *Day) findBadge(group *Group) int {
	badge := 0

	for item := range group.Bags[0].AllItems {
		_, inBagTwo := group.Bags[1].AllItems[item]
		_, inBagThree := group.Bags[2].AllItems[item]

		if inBagTwo && inBagThree {
			group.badge = item
			badge = d.PriorityMap[item]

			return badge
		}

	}

	return badge
}

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		PriorityMap: CreatePriorityMap(),
		InputFile:   "./day3/input.txt",
		GroupSize:   3,
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 3
}
