package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"path"
	"github.com/extrame/xls"
	"errors"
	"encoding/csv"
	"os"
	"github.com/axgle/mahonia"
)

//xlsx 和 xls只处理单sheet
func dealXlsx(filePath string) ([][]string, error) {
	data :=make ([][]string, 0)
	xlFile, err := xlsx.OpenFile(filePath)

	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return data, err
	}

	data = make ([][]string, len(xlFile.Sheets[0].Rows))
	for rIndex, row := range xlFile.Sheets[0].Rows {
		data[rIndex] = make([]string, len(row.Cells))
		for cIndex, cell := range row.Cells {
			text := cell.String()
			data[rIndex][cIndex]= text
		}
	}

	return data, nil
}

func dealXls(filePath string) ([][]string, error) {
	data := make ([][]string, 0)
	var xlFile *xls.WorkBook
	var err error

	if xlFile, err = xls.Open(filePath, "utf-8"); err != nil {
		return data, err
	}

	sheet1 := xlFile.GetSheet(0);

	if sheet1 == nil {
		return data, errors.New("have not sheet")
	}

	if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
		data = make ([][]string, int(sheet1.MaxRow)+1)
		for i := 0; i <= (int(sheet1.MaxRow)); i++ {
			row := sheet1.Row(i)
			data[i] = make([]string, row.LastCol()-row.FirstCol())
			first := row.FirstCol()
			last := row.LastCol()
			for j:=first ; j<last;j++  {
				col := row.Col(j)
				data[i][j]= col
			}
		}
	}

	return data, nil
}

func dealCsv(csvPath string) ([][]string, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	// r := csv.NewReader(file)
	r := csv.NewReader(decoder.NewReader(file))  // 这样，最终返回的字符串就是utf-8了。（go只认utf8）
	res, err := r.ReadAll()
	if err != nil {
		return nil,err
	}
	return res, nil
}

func DealExcel(filePath string) ([][]string, error) {
	var data [][]string
	var err error
	filenameWithSuffix := path.Base(filePath)
	fileSuffix := path.Ext(filenameWithSuffix) //获取文件后缀

	switch fileSuffix {
	case ".xlsx":
		data, err = dealXlsx(filePath)

	case ".xls":
		data, err = dealXls(filePath)

	case ".csv":
		data, err = dealCsv(filePath)
	}

	return data, err
}

func test1()  {
	//arr := [][]string
	arr := make([][]string, 10)

	fmt.Println(len(arr))

	var i int

	for i, _ := range arr{
		arr[i] = make([]string, 2)
		arr[i][0] ="123"
		arr[i][1] ="456"
		fmt.Println("i:%d", i)
	}

	//fmt.Println(arr)

	fmt.Println("==========")

	fmt.Println(i)
}


func main() {
	var data[][]string

	for i:=0; i<100000; i++ {
		data,_ = DealExcel("/Users/ybx/www/go_www2/go_study2/3.excel/1.xlsx")
		fmt.Println(data)
		data,_ = DealExcel("/Users/ybx/www/go_www2/go_study2/3.excel/1.xls")
		fmt.Println(data)
		data,_ = DealExcel("/Users/ybx/www/go_www2/go_study2/3.excel/1.csv")
		fmt.Println(data)
	}
}