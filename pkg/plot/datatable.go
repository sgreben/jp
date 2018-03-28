package plot

type DataTable struct {
	Columns []string
	Rows    [][]float64
}

func (d *DataTable) AddColumn(name string) {
	d.Columns = append(d.Columns, name)
}

func (d *DataTable) AddRow(elms ...float64) {
	d.Rows = append(d.Rows, elms)
}
