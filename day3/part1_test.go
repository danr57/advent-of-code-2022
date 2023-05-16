package day3_test

import (
	"testing"

	"github/danr57/advent-of-code-2022/day3"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	day := &day3.Day{
		PriorityMap: day3.CreatePriorityMap(),
		InputFile:   "example.txt",
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up day: %s", err)
	}

	result, err := day.Part1()
	if err != nil {
		t.Fatalf("error running part 1: %s", err)
	}

	if exampleAns := "157"; result != exampleAns {
		t.Fatalf("expected result to be %s, got %s", exampleAns, result)
	}
}
