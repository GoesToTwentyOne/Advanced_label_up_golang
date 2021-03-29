package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet2", "A5", "Myname is Nihad.")
	f.SetCellValue("Sheet1", "B5", 300)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("test.xlsx"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello this is run")

}
