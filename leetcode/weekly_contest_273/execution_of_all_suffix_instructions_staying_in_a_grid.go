package weekly_contest_273

type Position struct {
	x, y int
}

func (position Position) toNeighbour(direction byte) Position {
	switch direction {
	case 'U':
		position.x -= 1
	case 'D':
		position.x += 1
	case 'L':
		position.y -= 1
	case 'R':
		position.y += 1
	default:
		panic("Invalid Direction Prefix")
	}
	return position
}

func (position *Position) isInside(maxBounds *Position) bool {
	return position.x >= 0 && position.y >= 0 && position.x < maxBounds.x && position.y < maxBounds.y
}

func executeInstructions(n int, startPos []int, s string) []int {
	maxBounds, results := Position{x: n, y: n}, []int{}

	for i := 0; i < len(s); i++ {
		position, j := Position{x: startPos[0], y: startPos[1]}, i
		for ; j < len(s); j++ {
			position = position.toNeighbour(s[j])
			if !position.isInside(&maxBounds) {
				break
			}
		}
		results = append(results, j-i)
	}
	return results
}
