package util

import "testing"

func TestExcelRead(t *testing.T) {
	excelFile := "demo.xlsx"
	sheetIndex := 0
	data, err := ReadExcelData(excelFile, sheetIndex)
	if err != nil {
		t.Error(err)
	}
	for index, row := range data {
		t.Logf("row %d: %v", index, row)
	}
}
