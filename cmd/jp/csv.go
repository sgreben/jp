package main

import "strconv"

func parseCell(cell string) interface{} {
	f, err := strconv.ParseFloat(cell, 64)
	if err == nil {
		return f
	}
	b, err := strconv.ParseBool(cell)
	if err == nil {
		if b {
			return 1
		}
		return 0
	}
	return cell
}

func parseRows(rows [][]string) (out [][]interface{}) {
	out = make([][]interface{}, len(rows))
	for i, row := range rows {
		out[i] = make([]interface{}, len(row))
		for j, cell := range row {
			out[i][j] = parseCell(cell)
		}
	}
	return
}
