package day2

type (
	// Round contains the string value for our move and our opponent's move.
	Round struct {
		OpponentHand string
		MyHand       string
	}
)

const (
	loss     = 0
	draw     = 3
	win      = 6
	rock     = 1
	paper    = 2
	scissors = 3
)

// Reads in the Char to determine which move it represents
// 1 for rock, 2 for paper, 3 for scissors.
func convertMoveToValue(move string) int {
	switch move {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		return 0
	}
}

func (r *Round) part1() int {
	myHand := convertMoveToValue(r.MyHand)
	handToBeat := convertMoveToValue(r.OpponentHand)

	switch myHand - handToBeat {
	case -1, 2: //nolint:gomnd // This is a puzzle
		return loss + myHand
	case 0:
		return draw + myHand
	case 1, -2:
		return win + myHand
	default:
		return 0
	}
}

func (r *Round) part2() int {
	opponentHand := convertMoveToValue(r.OpponentHand)

	switch r.MyHand {
	case "X": // we need to lose
		if opponentHand == rock {
			return loss + scissors
		}

		return loss + opponentHand - 1

	case "Y": // we need to draw
		return draw + opponentHand
	case "Z": // we need to win
		if opponentHand == scissors {
			return win + rock
		}

		return win + opponentHand + 1

	default:
		return 0
	}
}
