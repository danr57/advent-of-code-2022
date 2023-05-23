package day6

import (
	"fmt"
	"os"
	"strings"
)

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	data, err := os.ReadFile("day6/input.txt")
	if err != nil {
		return fmt.Errorf("couldn't open input file: %w", err)
	}

	d.DataStream = strings.Split(string(data), "")

	return nil
}
