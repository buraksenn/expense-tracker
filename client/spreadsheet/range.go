package spreadsheet

import "fmt"

type Range struct {
	SheetName string
	Row       int
	Col       int
}

func (r *Range) String() string {
	return fmt.Sprintf("%s!%d:%d", r.SheetName, r.Col, r.Row)
}
