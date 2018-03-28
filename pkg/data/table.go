package data

type Table struct {
	Columns []string
	Rows    [][]float64
}

func (d *Table) AddColumn(name string) {
	d.Columns = append(d.Columns, name)
}

func (d *Table) AddRow(elms ...float64) {
	d.Rows = append(d.Rows, elms)
}
