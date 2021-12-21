package grid

// Interface used as a generic
type T interface{}

// Grid is a dwo dimentional array that can be walked
// Left(), Right(), Up(), Down() cell by cell infinitely
type Grid struct {
	cells []T
	dim   int
}

// Make new Grid by providing length of one side, Grid will be dim x dim
// If dim = 3, Grid will be 3 by 3
func NewGrid(dim int) *Grid {
	g := Grid{
		cells: make([]T, dim*dim),
		dim:   dim,
	}
	for i := range g.cells {
		g.cells[i] = 0
	}
	return &g
}

// Make new Grid by providing length of one side, Grid will be dim x dim
// If dim = 3, Grid will be 3 by 3
func New(dim int) *Grid {
	return NewGrid(dim)
}

// Public getter to get all cells as a single array,
func (g Grid) Cells() []T {
	return g.cells
}

// Get one dimension size
func (g Grid) Dim() int {
	return g.dim
}

// Get index of a Row by providing Index of cell
func (g Grid) Row(i int) int {
	return int(i / g.dim)
}

// Get index of a column by providing Index of a cell
func (g Grid) Column(i int) int {
	return i % g.dim
}

// Get index and a Row and Column by providing Index of a cell
func (g Grid) RowColumn(i int) (int, int) {
	return g.Row(i), g.Column(i)
}

// Get value of a cell by providing Index of a cell
func (g Grid) Value(i int) T {
	return g.cells[i]
}

// Get Index of a cell by providing row and column index
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

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Up(i int) int {
	if g.isFirstRow(i) {
		return g.Index(g.dim-1, g.Column(i))
	}
	return g.Index(g.Row(i)-1, g.Column(i))
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Down(i int) int {
	if g.isLastRow(i) {
		return g.Index(0, g.Column(i))
	}
	return g.Index(g.Row(i)+1, g.Column(i))
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Left(i int) int {
	if g.isFirstColumn(i) {
		return g.Index(g.Row(i), g.dim-1)
	}
	return g.Index(g.Row(i), g.Column(i)-1)
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Right(i int) int {
	if g.isLastColumn(i) {
		return g.Index(g.Row(i), 0)
	}
	return g.Index(g.Row(i), g.Column(i)+1)
}

// Set value of a cell where Index is first arg and value second
func (g Grid) Set(i int, v int) {
	g.cells[i] = v
}
