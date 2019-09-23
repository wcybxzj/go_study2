package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	row = sheet.AddRow()
	//row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "mobile"
	cell = row.AddCell()
	cell.Value = "userid"

	for i:=0;i < 100000 ;i++  {
		row = sheet.AddRow()
		//row.SetHeightCM(1)
		cell = row.AddCell()
		cell.Value = "17710750086"
		/*
		cell = row.AddCell()
		cell.Value = "0"
		*/
	}

	err = file.Save("test_write.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}