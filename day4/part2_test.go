package day4_test

import (
	"github/danr57/advent-of-code-2022/day4"
	"testing"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	day := &day4.Day{
		InputFile: "example.txt",
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up day: %s", err)
	}

	result, err := day.Part2()
	if err != nil {
		t.Fatalf("error running part 2: %s", err)
	}

	if exampleAns := "4"; result != exampleAns {
		t.Fatalf("expected result to be %s, got %s", exampleAns, result)
	}
}
