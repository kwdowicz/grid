package grid

type Grid struct {
	cells []int
	dim   int
}

func NewGrid(dim int) *Grid {
	g := Grid{
		cells: make([]int, dim*dim),
		dim:   dim,
	}
	return &g
}

func (g Grid) Cells() []int {
	return g.cells
}

func (g Grid) Dim() int {
	return g.dim
}

func (g Grid) Row(i int) int {
	return int(i / g.dim)
}

func (g Grid) Column(i int) int {
	return i % g.dim
}

func (g Grid) RowColumn(i int) (int, int) {
	return g.Row(i), g.Column(i)
}

func (g Grid) Value(i int) int {
	return g.cells[i]
}

func (g Grid) Index(r int, c int) int {
	return (g.dim * r) + c
}

func (g Grid) isFirstRow(i int) bool {
	return g.Row(i) == 0
}

func (g Grid) isLastRow(i int) bool {
	return g.Row(i) == (g.dim - 1)
}

func (g Grid) isFirstColumn(i int) bool {
	return g.Column(i) == 0
}

func (g Grid) isLastColumn(i int) bool {
	return g.Column(i) == (g.dim - 1)
}

func (g Grid) Up(i int) int {
	if g.isFirstRow(i) {
		return g.Index(g.dim-1, g.Column(i))
	}
	return g.Index(g.Row(i)-1, g.Column(i))
}

func (g Grid) Down(i int) int {
	if g.isLastRow(i) {
		return g.Index(0, g.Column(i))
	}
	return g.Index(g.Row(i)+1, g.Column(i))
}

func (g Grid) Left(i int) int {
	if g.isFirstColumn(i) {
		return g.Index(g.Row(i), g.dim-1)
	}
	return g.Index(g.Row(i), g.Column(i)-1)
}

func (g Grid) Right(i int) int {
	if g.isLastColumn(i) {
		return g.Index(g.Row(i), 0)
	}
	return g.Index(g.Row(i), g.Column(i)+1)
}

func (g Grid) Set(i int, v int) {
	g.cells[i] = v
}

func (g Grid) SumAround(i int) int {
	return g.Value(g.Up(i)) + g.Value(g.Right(g.Up(i))) +
		g.Value(g.Right(i)) + g.Value(g.Right(g.Down(i))) +
		g.Value(g.Down(i)) + g.Value(g.Down(g.Left(i))) +
		g.Value(g.Left(i)) + g.Value(g.Left(g.Up(i)))
}
