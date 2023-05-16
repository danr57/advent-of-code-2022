package day2

type (
	Round struct {
		OpponentHand string
		MyHand       string
		Score        int
	}
)

const (
	loss = 0
	draw = 3
	win  = 6
)

// PlayRound will determine and return the score from a given round
func (r *Round) playRound() int {

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
