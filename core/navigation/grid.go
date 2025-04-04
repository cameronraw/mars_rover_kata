package navigation

type Grid struct {
	Tiles [][]int
}

func NewGrid(width, height int) *Grid {
	grid := &Grid{
		Tiles: make([][]int, height),
	}
	for i := range grid.Tiles {
		grid.Tiles[i] = make([]int, width)
	}
	return grid
}
