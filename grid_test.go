package grid

import "testing"

func TestRow1(t *testing.T) {
	g := NewGrid(3)

	r := g.Row(1)
	if r != 0 {
		t.Errorf("Row of 1 should be 0 in dim 3 grid, but is: %d ", r)
	}
}

func TestRow17(t *testing.T) {
	g := NewGrid(5)

	r := g.Row(17)
	if r != 3 {
		t.Errorf("Row of 17 should be 3 in dim 5 grid, but is: %d ", r)
	}
}

func TestColumn17(t *testing.T) {
	g := NewGrid(5)

	r := g.Column(17)
	if r != 2 {
		t.Errorf("Column of 17 should be 2 in dim 5 grid, but is: %d ", r)
	}
}

func TestColumn6(t *testing.T) {
	g := NewGrid(3)

	r := g.Column(6)
	if r != 0 {
		t.Errorf("Column of 6 should be 0 in dim 3 grid, but is: %d ", r)
	}
}

func TestRowColumn(t *testing.T) {
	g := NewGrid(5)

	r, c := g.RowColumn(11)

	if r != 2 || c != 1 {
		t.Errorf("Row should be %d and is %d, Column is %d and is %d", 2, r, 1, c)
	}
}

func TestFirstRow(t *testing.T) {
	g := NewGrid(3)

	if !g.isFirstRow(1) {
		t.Error("1 in grid 3 should be first row")
	}
}

func TestLastRow(t *testing.T) {
	g := NewGrid(3)

	if !g.isLastRow(7) {
		t.Error("7 in grid 3 should be last row")
	}
}

func TestFirstColumn(t *testing.T) {
	g := NewGrid(3)

	if !g.isFirstColumn(3) {
		t.Error("3 in grid 3 should be first column")
	}
}

func TestLastColumn(t *testing.T) {
	g := NewGrid(3)

	if !g.isLastColumn(5) {
		t.Error("5 in grid 3 should be last row")
	}
}

func TestValue(t *testing.T) {
	g := NewGrid(5)
	g.cells[17] = 99
	r := g.Value(17)
	if r != 99 {
		t.Errorf("Result should be 99 not: %d ", r)
	}
}
func TestGet(t *testing.T) {
	g := NewGrid(5)
	g.cells[17] = 99
	r := g.Value(g.Index(3, 2))
	if r != 99 {
		t.Errorf("Result should be 99 not: %d ", r)
	}
}

func TestUp(t *testing.T) {
	g := NewGrid(3)
	g.cells[1] = 33
	g.cells[8] = 11
	r := g.Value(g.Up(4))
	z := g.Value(g.Up(2))
	if r != 33 || z != 11 {
		t.Errorf("Result should be r=33, z=11, not: %d and %d", r, z)
	}
}
func TestDown(t *testing.T) {
	g := NewGrid(3)
	g.cells[1] = 33
	g.cells[8] = 11
	r := g.Value(g.Down(7))
	z := g.Value(g.Down(5))
	if r != 33 || z != 11 {
		t.Errorf("Result should be r=33, z=11, not: %d and %d", r, z)
	}
}

func TestLeft(t *testing.T) {
	g := NewGrid(3)
	g.cells[1] = 33
	g.cells[8] = 11
	r := g.Value(g.Left(2))
	z := g.Value(g.Left(6))
	if r != 33 || z != 11 {
		t.Errorf("Result should be r=33, z=11, not: %d and %d", r, z)
	}
}
func TestRight(t *testing.T) {
	g := NewGrid(3)
	g.cells[0] = 33
	g.cells[8] = 11
	r := g.Value(g.Right(2))
	z := g.Value(g.Right(7))
	if r != 33 || z != 11 {
		t.Errorf("Result should be r=33, z=11, not: %d and %d", r, z)
	}
}

func TestSumAround(t *testing.T) {
	g := NewGrid(3)

	for i := range g.cells {
		g.Set(i, i)
	}

	if g.SumAround(4) != 32 {
		t.Errorf("Sum Around should be 32 but is %d", g.SumAround(4))
	}
}

func TestSet(t *testing.T) {
	g := NewGrid(3)
	g.Set(g.Index(2, 2), 33)
	r := g.Value(8)
	if r != 33 {
		t.Errorf("Index 8 is %d, instead of 33", r)
	}
}
