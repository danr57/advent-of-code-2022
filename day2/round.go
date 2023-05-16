package day2

type (
	Round struct {
		OpponentHand string
		MyHand       string
	}
)

const (
	loss = 0
	draw = 3
	win  = 6
)

// Reads in the Char to determine which move it represents
// 1 for rock, 2 for paper, 3 for scissors
func convertMoveToValue(move string) int {
	switch move {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		return 0
	}
}

func (r *Round) part1() int {
	myHand := convertMoveToValue(r.MyHand)
	handToBeat := convertMoveToValue(r.OpponentHand)

	switch myHand - handToBeat {
	case -1, 2:
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
		if opponentHand == 1 {
			return loss + 3
		}
		return loss + opponentHand - 1
	case "Y": // we need to draw
		return draw + opponentHand
	case "Z": // we need to win
		if opponentHand == 3 {
			return win + 1
		}
		return win + opponentHand + 1
	default:
		return 0
	}
}
