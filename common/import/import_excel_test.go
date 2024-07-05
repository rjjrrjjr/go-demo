package _import

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestImportExcelTealeg(t *testing.T) {
	file, err := xlsx.OpenFile("../export/sample.xlsx")
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	sheet := file.Sheets[0]
	for i := 0; i < sheet.MaxRow; i++ {
		for j := 0; j < sheet.MaxCol; j++ {
			cell, _ := sheet.Cell(i, j)
			fmt.Printf("%s\t", cell.String())
		}
		fmt.Println()
	}
}

func TestImportExcelPopularExcelize(t *testing.T) {
	f, err := excelize.OpenFile("../export/sample.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.SheetCount)
	fmt.Println(f.Sheet)

	// 读取第一个工作表中的所有行
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历每一行并打印数据
	for _, row := range rows {
		for _, cell := range row {
			fmt.Printf("%s\t", cell)
		}
		fmt.Println()
	}
}
