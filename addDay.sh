#!/bin/bash

max=$(find . -type d  | grep -E 'day[0-9]+' | cut -c6- | sort -r | head -n 1)

next=$(printf 'day%d' "$((max + 1))")


mkdir -p "$next"

touch "$next"/input.txt
touch "$next"/example.txt
touch "$next"/readme.md


cat > "$next"/day.go << EOM
// Package $next implements the solution to Day $((max+1)) of the Advent of Code 2022.
package $next

type (
	// Day is the implementation of Day $((max+1)).
	Day struct{
		InputFile string
	}
)

// New returns a new instance of Day.
func New() *Day {
	return &Day{}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return $((max + 1))
}

EOM

cat > "$next"/setup.go << EOM
package $next

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	return nil
}

EOM


cat > "$next"/part1.go << EOM
package $next

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	return "", nil
}

EOM

cat > "$next"/part1_test.go << EOM
package ${next}_test

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()
}

EOM

cat > "$next"/part2.go << EOM
package $next

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	return "", nil
}

EOM

cat > "$next"/part2_test.go << EOM
package ${next}_test

import (
	"testing"
)

func TestPart2(t *testing.T) {
	t.Parallel()
}

EOM



