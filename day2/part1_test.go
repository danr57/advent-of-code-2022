package day2_test

import (
	"github/danr57/advent-of-code-2022/day2"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	exampleAns := "15"

	day := &day2.Day{
		InputFile: "example.txt",
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up day: %s", err)
	}

	result, err := day.Part1()
	if err != nil {
		t.Fatalf("error running part 1: %s", err)
	}

	if result != exampleAns {
		t.Fatalf("expected result to be %s, got %s", exampleAns, result)
	}
}
