package table

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// TableWriter will write the given row data
// in form of a pretty table
type TableWriter struct {
	out           io.Writer
	ColumnHeaders []string
	ColumnWidths  []int
}

// AlignType is type of alignment to use for when logging
type AlignType int

const (
	// AlignCenter will align given text to center of space available
	AlignCenter AlignType = iota
	// AlignLeft will align given text to left
	AlignLeft
	// AlignRight will align given text to right
	AlignRight
)

// NewTableWriter will initialize the table writer to format data
func NewTableWriter(columnHeaders []string, columnWidths []int) (*TableWriter, error) {
	return NewTableWriterWithWriter(os.Stdout, columnHeaders, columnWidths)
}

// NewTableWriterWithWriter will initialize the table writer with given writer to format data
func NewTableWriterWithWriter(writer io.Writer, columnHeaders []string, columnWidths []int) (*TableWriter, error) {
	tw := TableWriter{
		writer,
		columnHeaders,
		columnWidths,
	}

	err := tw.validate()
	if err != nil {
		return nil, err
	}
	return &tw, nil
}

// validate will validate if the column widths provided are at least
// the size of column headers
func (t *TableWriter) validate() error {
	for i, h := range t.ColumnHeaders {
		if len(h) > t.ColumnWidths[i] {
			return fmt.Errorf("Header length can't be larger than width's provided")
		}
	}
	return nil
}

// PrintHeader will print the header of the table with column headers
func (t *TableWriter) PrintHeader() {
	t.PrintFooter()
	t.PrintRow(t.ColumnHeaders, AlignCenter)
	t.PrintFooter()
}

// PrintRow will print data given in form a row in a table
func (t *TableWriter) PrintRow(row []string, alignmentType AlignType) error {
	for i, v := range row {
		if len(v) > t.ColumnWidths[i] {
			return fmt.Errorf("Column value '%s' is longer than max width '%d' specified", v, t.ColumnWidths[i])
		}
		if alignmentType == 1 {
			row[i] = t.alignRight(v, t.ColumnWidths[i])
		} else if alignmentType == 2 {
			row[i] = t.alignLeft(v, t.ColumnWidths[i])
		} else {
			row[i] = t.alignCenter(v, t.ColumnWidths[i])
		}
	}
	t.out.Write([]byte(t.repeatWithValue("|", row) + "|\n"))
	return nil
}

// PrintRows will print given rows only
func (t *TableWriter) PrintRows(rows [][]string, alignmentType AlignType) error {
	for _, row := range rows {
		err := t.PrintRow(row, alignmentType)
		if err != nil {
			return err
		}
	}
	return nil
}

// PrintRowAsOneColumn prints given value as single column
func (t *TableWriter) PrintRowAsOneColumn(value string, alignmentType AlignType) error {
	totalWidth := len(t.ColumnHeaders) - 1
	for _, w := range t.ColumnWidths {
		totalWidth = totalWidth + w
	}
	if len(value) > totalWidth {
		return fmt.Errorf("Column value '%s' is longer than max width '%d' specified", value, totalWidth)
	}
	if alignmentType == 1 {
		t.out.Write([]byte("|" + t.alignRight(value, totalWidth) + "|\n"))
	} else if alignmentType == 2 {
		t.out.Write([]byte("|" + t.alignLeft(value, totalWidth) + "|\n"))
	} else {
		t.out.Write([]byte("|" + t.alignCenter(value, totalWidth) + "|\n"))
	}
	return nil
}

// PrintFooter will print footer to close the table
func (t *TableWriter) PrintFooter() {
	t.out.Write([]byte(t.decorator()))
}

// PrintTable will print given rows data as table at once
func (t *TableWriter) PrintTable(rows [][]string, alignmentType AlignType) {
	t.PrintHeader()
	t.PrintRows(rows, alignmentType)
	t.PrintFooter()
}

func (t *TableWriter) repeatWithValue(s string, values []string) string {
	out := ""
	for _, v := range values {
		out = out + s + v
	}
	return out
}

func (t *TableWriter) decorator() string {
	s := ""
	for i := range t.ColumnHeaders {
		s = fmt.Sprintf("%s%s%s", s, "+", strings.Repeat("-", t.ColumnWidths[i]))
	}
	return s + "+\n"
}

func (t *TableWriter) alignCenter(s string, width int) string {
	remainingChars := width - len(s)
	if remainingChars > 0 {
		leftPadSize := int(math.Ceil(float64(remainingChars / 2)))
		rightPadSize := remainingChars - leftPadSize
		return strings.Repeat(" ", leftPadSize) + s + strings.Repeat(" ", rightPadSize)
	}
	return s
}

func (t *TableWriter) alignRight(s string, width int) string {
	remainingChars := width - len(s)
	if remainingChars > 0 {
		return s + strings.Repeat(" ", remainingChars)
	}
	return s
}

func (t *TableWriter) alignLeft(s string, width int) string {
	remainingChars := width - len(s)
	if remainingChars > 0 {
		return strings.Repeat(" ", remainingChars) + s
	}
	return s
}
