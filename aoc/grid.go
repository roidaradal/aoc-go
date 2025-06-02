package aoc

var (
	U Delta = [2]int{-1, 0}
	D Delta = [2]int{1, 0}
	L Delta = [2]int{0, -1}
	R Delta = [2]int{0, 1}
)

func Move(c Coords, d Delta) Coords {
	row, col := c[0], c[1]
	dy, dx := d[0], d[1]
	return [2]int{row + dy, col + dx}
}
